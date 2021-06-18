package cws

import (
	"fmt"

	"im-server/protoim"
	"im-server/servers"
)

// GetUserRequest 获取用户信息-请求
func GetUserRequest(c *servers.Context) {
	req := new(protoim.GetUserRequest)
	err := c.ProtoBind(req)
	if err != nil {
		c.SendError(protoim.Error_E_PARAMETER, err.Error())
		return
	}

	fmt.Println("GetUserRequest - req:", req)

	rsp := &protoim.GetUserResponse{
		User: &protoim.UserInfo{
			UserId:   req.UserId,
			NickName: "laixhe",
			Age:      18,
			IsOk:     true,
		},
	}

	e := c.Send(protoim.CMD_C_GET_USER_RESPONSE, rsp)
	fmt.Println("GetUserInfo - rsp:", rsp, "err:", e)
}
