package cws

import (
	"fmt"

	"im-server/protoim"
	"im-server/servers"
)

// GetUserInfoRequest 获取用户信息-请求
func GetUserInfoRequest(c *servers.Context) {
	req := new(protoim.GetUserInfoRequest)
	err := c.ProtoBind(req)
	if err != nil {
		c.SendError(protoim.E_PARAMETER_ERROR, err.Error())
		return
	}

	rsp := &protoim.GetUserInfoResponse{
		User: &protoim.UserInfo{
			Uid:  req.Uid,
			Name: "laixhe",
			Age:  18,
			IsOk: true,
		},
	}

	e := c.Send(protoim.CMD_GET_USER_INFO_RESPONSE, rsp)
	fmt.Println("GetUserInfo - rsp:", rsp, "err:", e)
}
