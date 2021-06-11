package servers

import (
	"im-server/protoim"
)

// ping 心跳
func ping(c *Context) {
	// 响应 pong
	c.SendCmd(protoim.CMD_C_PONG)
}
