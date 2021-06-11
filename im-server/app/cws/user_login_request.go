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
		c.SendError(protoim.Error_E_PARAMETER, err.Error())
		return
	}

	if !utils.IsAppOs(req.AppOsId) {
		c.SendError(protoim.Error_E_PARAMETER, "app os error")
		return
	}

	fmt.Println("UserLogin - req:", req)
	rsp := &protoim.UserLoginResponse{
		UserId:   cryptos.Md5(req.Account),
		NickName: basictype.GetRandomString(6, 1),
	}

	e := c.Send(protoim.CMD_C_USER_LOGIN_RESPONSE, rsp)
	if e == nil {
		c.UserLogin(req.AppOsId, rsp.UserId)
	}

	fmt.Println("UserLogin - rsp:", rsp)
	c.Sendxxx(protoim.CMD_C_UPDATE_FRIEND_INFO, &protoim.UpdateFriendInfo{
		Tag: "add",
		User: &protoim.UserInfo{
			UserId:   rsp.UserId,
			NickName: rsp.NickName,
		},
	})
}
