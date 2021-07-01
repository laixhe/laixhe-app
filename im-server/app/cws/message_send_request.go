package cws

import (
	"fmt"
	"im-server/protoim"
	"im-server/servers"
)

// MessageSendRequest 消息相关(客户端发送-服务端接收)-请求
func MessageSendRequest(c *servers.Context) {
	req := new(protoim.MessageSendRequest)
	err := c.ProtoBind(req)
	if err != nil {
		c.SendError(protoim.ErrorType_E_PARAMETER, err.Error())
		return
	}

	fmt.Println("MessageRequest - req:", req)

	rsp := &protoim.MessageFeedback{
	}

	e := c.Send(protoim.CMD_C_MESSAGE_FEEDBACK, rsp)
	fmt.Println("MessageRequest - rsp:", rsp, "err:", e)
}
