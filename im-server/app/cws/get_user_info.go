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
		fmt.Println("err:", err)
		return
	}
	rsp := &protoim.User{
		Uid:  req.Uid,
		Name: "laixhe",
		Age:  18,
		Isok: true,
	}
	err = c.Send(protoim.CMD_CHECKING_TOKEN, rsp)
	fmt.Println("rsp:", rsp, "err:", err)
}
