package servers

import "im-server/protoim"

// Context 客户端请求的数据上下文
type Context struct {
	conn *Client     // 客户端链接
	msg  []byte      // 客户端请求的数据
	cmd  protoim.CMD // 客户端请求指令
}

// NewContext init
func NewContext(conn *Client, msg []byte, cmd protoim.CMD) *Context {
	return &Context{
		conn: conn,
		msg:  msg,
		cmd:  cmd,
	}
}
