package routers

import (
	"im-server/app/cws"
	"im-server/protoim"
	"im-server/servers"
)

// initWs 初始化业务路由
func initWs(r *servers.Router) {
	r.Set(protoim.CMD_GET_USER_INFO, cws.GetUserInfo)
}
