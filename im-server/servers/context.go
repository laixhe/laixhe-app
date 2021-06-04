package servers

import (
	"github.com/laixhe/goutil/zaplog"
	"google.golang.org/protobuf/proto"

	"im-server/protoim"
)

// Context 客户端请求的数据上下文
type Context struct {
	conn *Client     // 客户端链接
	data []byte      // 客户端请求的数据
	cmd  protoim.CMD // 客户端请求指令
}

// NewContext init
func NewContext(conn *Client, data []byte, cmd protoim.CMD) *Context {
	return &Context{
		conn: conn,
		data: data,
		cmd:  cmd,
	}
}

// ProtoBind 数据绑定
func (c *Context) ProtoBind(data proto.Message) error {
	return proto.Unmarshal(c.data, data)
}

// Send 发送数据
func (c *Context) Send(cmd protoim.CMD, data proto.Message) *protoim.ErrorInfo {
	// 判断当前链接是否已经关闭
	if c.conn.isClosed {
		return nil
	}
	protoBase, err := EnCode(cmd, data)
	if err != nil {
		return ErrorMessage(protoim.E_ENCODE_ERROR, err.Error())
	}
	// 捕获 panic (注要是发送消息时可能通道已关闭)
	defer func() {
		if err := recover(); err != nil {
			zaplog.Errorf("context send=%v err=%v", c.conn.addr, err)
		}
	}()

	c.conn.send <- protoBase
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
		return ErrorMessage(protoim.E_ENCODE_ERROR, err.Error())
	}
	// 捕获 panic (注要是发送消息时可能通道已关闭)
	defer func() {
		if err := recover(); err != nil {
			zaplog.Errorf("context send=%v err=%v", c.conn.addr, err)
		}
	}()

	c.conn.send <- protoBase
	return nil
}

func (c *Context) Sendxxx(cmd protoim.CMD, data proto.Message) *protoim.ErrorInfo {
	// 判断当前链接是否已经关闭
	if c.conn.isClosed {
		return nil
	}
	protoBase, err := EnCode(cmd, data)
	if err != nil {
		return ErrorMessage(protoim.E_ENCODE_ERROR, err.Error())
	}
	// 捕获 panic (注要是发送消息时可能通道已关闭)
	defer func() {
		if err := recover(); err != nil {
			zaplog.Errorf("context send=%v err=%v", c.conn.addr, err)
		}
	}()

	c.conn.manager.broadcast <- protoBase
	return nil
}
