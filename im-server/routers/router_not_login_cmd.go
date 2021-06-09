package routers

import (
	"im-server/app/service"
	"im-server/protoim"
)

// notLoginCmd 添加不需要登录的路由就可以访问的指令
func notLoginCmd() {
	service.AddNotLoginCmd(protoim.CMD_USER_LOGIN_REQUEST)
}
