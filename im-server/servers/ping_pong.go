package servers

import (
	"im-server/protoim"
)

// Ping 心跳
func Ping(c *Context) {
	// 响应 pong
	c.SendCmd(protoim.CMD_PONG)
}
