package cws

import (
	"fmt"

	"im-server/protoim"
	"im-server/servers"
	"im-server/utils"
)

// UserLoginRequest 用户登录-请求
func UserLoginRequest(c *servers.Context) {
	req := new(protoim.UserLoginRequest)
	err := c.ProtoBind(req)
	if err != nil {
		c.SendError(protoim.ErrorType_E_PARAMETER, err.Error())
		return
	}

	if !utils.IsAppOs(req.OsId) {
		c.SendError(protoim.ErrorType_E_PARAMETER, "app os error")
		return
	}

	fmt.Println("UserLoginRequest - req:", req)

	rsp := &protoim.UserLoginResponse{
	}

	e := c.Send(protoim.CMD_C_USER_LOGIN_RESPONSE, rsp)
	if e == nil {
	}

}
