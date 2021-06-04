package cws

import (
	"fmt"

	"im-server/protoim"
	"im-server/servers"
)

// UserLogin 用户登录
func UserLogin(c *servers.Context) {
	req := new(protoim.UserLoginRequest)
	err := c.ProtoBind(req)
	if err != nil {
		fmt.Println("UserLogin - err:", err)
		return
	}

	rsp := &protoim.UserLoginResponse{
		Uid:  req.Uid,
		Name: req.Name,
	}

	e := c.Send(protoim.CMD_USER_LOGIN_RESPONSE, rsp)
	fmt.Println("UserLogin - rsp:", rsp, "err:", e)

	c.Sendxxx(protoim.CMD_UPDATE_USER_FRIEND_INFO, &protoim.UpdateUserFriendInfo{
		Tag:  "add",
		Uid:  req.Uid,
		Name: req.Name,
	})
}
