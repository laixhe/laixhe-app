package cws

import (
	"fmt"

	"github.com/laixhe/goutil/basictype"
	"github.com/laixhe/goutil/cryptos"

	"im-server/protoim"
	"im-server/servers"
	"im-server/utils"
)

// UserLoginRequest 用户登录-请求
func UserLoginRequest(c *servers.Context) {
	req := new(protoim.UserLoginRequest)
	err := c.ProtoBind(req)
	if err != nil {
		c.SendError(protoim.E_PARAMETER_ERROR, err.Error())
		return
	}

	if utils.IsAppOs(req.AppOsId) {
		c.SendError(protoim.E_PARAMETER_ERROR, "app os error")
		return
	}

	fmt.Println("UserLogin - req:", req)
	rsp := &protoim.UserLoginResponse{
		Uid:  cryptos.Md5(req.Account),
		Name: basictype.GetRandomString(6, 1),
	}

	e := c.Send(protoim.CMD_USER_LOGIN_RESPONSE, rsp)
	if e == nil {
		c.UserLogin(req.AppOsId, rsp.Uid)
	}

	fmt.Println("UserLogin - rsp:", rsp)
	c.Sendxxx(protoim.CMD_UPDATE_USER_FRIEND_INFO, &protoim.UpdateUserFriendInfo{
		Tag:  "add",
		Uid:  rsp.Uid,
		Name: rsp.Name,
	})
}
