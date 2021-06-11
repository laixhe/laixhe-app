package routers

import (
	"im-server/app/cws"
	"im-server/protoim"
	"im-server/servers"
)

// wsUser 初始化业务路由
func wsUser(r *servers.Router) {
	r.Set(protoim.CMD_C_USER_LOGIN_REQUEST, cws.UserLoginRequest)
	r.Set(protoim.CMD_C_GET_USER_INFO_REQUEST, cws.GetUserInfoRequest)
	r.Set(protoim.CMD_C_MESSAGE_REQUEST, cws.MessageRequest)
}
