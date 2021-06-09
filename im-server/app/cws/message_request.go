package cws

import (
	"fmt"
	"time"

	"github.com/laixhe/goutil/utils"

	"im-server/app/service"
	"im-server/protoim"
	"im-server/servers"
)

// MessageRequest 消息相关(客户端发送-服务端接收)-请求
func MessageRequest(c *servers.Context) {
	req := new(protoim.MessageRequest)
	err := c.ProtoBind(req)
	if err != nil {
		c.SendError(protoim.E_PARAMETER_ERROR, err.Error())
		return
	}
	fmt.Println("MessageRequest - req:", req)

	rsp := &protoim.MessageFeedback{
		LocalMsgId: req.LocalMsgId,
	}

	e := c.Send(protoim.CMD_MESSAGE_FEEDBACK, rsp)
	fmt.Println("MessageRequest - rsp:", rsp, "err:", e)

	data := &protoim.MessageResponse{
		MsgId:         fmt.Sprintf("%d", utils.GetXUID()),
		LocalMsgId:    req.LocalMsgId,
		Pts:           0,
		FromId:        req.FromId,
		ChatTypeId:    req.ChatTypeId,
		ToId:          req.ToId,
		MessageTypeId: req.MessageTypeId,
		Content:       req.Content,
		DataTime:      time.Now().Unix(),
	}

	_ = service.SendMessageResponse(data)
}
