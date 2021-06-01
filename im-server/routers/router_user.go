package routers

import (
	"im-server/app/cws"
	"im-server/protoim"
	"im-server/servers"
)

// wsUser 初始化业务路由
func wsUser(r *servers.Router) {
	r.Set(protoim.CMD_GetUserInfoReq, cws.GetUserInfo)
}
