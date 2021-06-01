package cws

import (
	"fmt"

	"im-server/protoim"
	"im-server/servers"
)

// GetUserInfo 获取用户信息
func GetUserInfo(c *servers.Context) {
	req := new(protoim.GetUserInfoRequest)
	err := c.ProtoBind(req)
	if err != nil {
		fmt.Println("GetUserInfo - err:", err)
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

	e := c.Send(protoim.CMD_GetUserInfoRes, rsp)
	fmt.Println("GetUserInfo - rsp:", rsp, "err:", e)
}
