package routers

import (
	"im-server/app/cws"
	"im-server/protoim"
	"im-server/servers"
)

// wsUser 初始化业务路由
func wsUser(r *servers.Router) {
	r.Set(protoim.CMD_USER_LOGIN_REQUEST, cws.UserLogin)
	r.Set(protoim.CMD_GET_USER_INFO_REQUEST, cws.GetUserInfo)
}
