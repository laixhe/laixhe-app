package servers

import (
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
func (c *Context) Send(cmd protoim.CMD, data proto.Message) error {
	// 判断当前链接是否已经关闭
	if c.conn.isClosed {
		return nil
	}
	protoData, err := proto.Marshal(data)
	if err != nil {
		return err
	}
	messageBase := &protoim.MessageBase{
		Cmd:  cmd,
		Data: protoData,
	}
	protoBase, err := proto.Marshal(messageBase)
	if err != nil {
		return err
	}
	c.conn.send <- protoBase
	return nil
}
