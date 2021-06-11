package servers

import (
	"github.com/laixhe/goutil/zaplog"
	"google.golang.org/protobuf/proto"

	"im-server/protoim"
)

// Context 客户端请求的数据上下文
type Context struct {
	conn *client     // 客户端链接
	data []byte      // 客户端请求的数据
	cmd  protoim.CMD // 客户端请求指令
}

// ProtoBind 数据绑定
func (c *Context) ProtoBind(data proto.Message) error {
	return proto.Unmarshal(c.data, data)
}

// UserLogin 用户登录
func (c *Context) UserLogin(appOsId protoim.AppOs, userId string) {
	c.conn.userLogin(appOsId, userId)
}

// Send 发送数据
func (c *Context) Send(cmd protoim.CMD, data proto.Message) *protoim.ErrorInfo {
	// 判断当前链接是否已经关闭
	if c.conn.isClosed {
		return nil
	}
	protoBase, err := EnCode(cmd, data)
	if err != nil {
		return ErrorMessage(protoim.Error_E_ENCODE, err.Error())
	}
	_ = c.conn.sendData(protoBase)
	return nil
}

// Send 发送数据 CMD
func (c *Context) SendCmd(cmd protoim.CMD) *protoim.ErrorInfo {
	// 判断当前链接是否已经关闭
	if c.conn.isClosed {
		return nil
	}
	protoBase, err := EnCodeCmd(cmd)
	if err != nil {
		return ErrorMessage(protoim.Error_E_ENCODE, err.Error())
	}
	_ = c.conn.sendData(protoBase)
	return nil
}

func (c *Context) Sendxxx(cmd protoim.CMD, data proto.Message) *protoim.ErrorInfo {
	// 判断当前链接是否已经关闭
	if c.conn.isClosed {
		return nil
	}
	protoBase, err := EnCode(cmd, data)
	if err != nil {
		return ErrorMessage(protoim.Error_E_ENCODE, err.Error())
	}
	// 捕获 panic (注要是发送消息时可能通道已关闭)
	defer func() {
		if err := recover(); err != nil {
			zaplog.Errorf("context-sendxxx addr=%v err=%v", c.conn.addr, err)
		}
	}()

	c.conn.manager.broadcast <- protoBase
	return nil
}

// Error 发送错误
func (c *Context) SendError(e protoim.Error, msg ...string) {
	// 判断当前链接是否已经关闭
	if c.conn.isClosed {
		return
	}
	protoBase, err := EnCodeError(e, msg...)
	if err != nil {
		zaplog.Errorf("context-senderror addr=%v err=%v", c.conn.addr, err)
		return
	}
	_ = c.conn.sendData(protoBase)
}
