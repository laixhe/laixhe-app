package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/laixhe/goutil/zaplog"

	"im-server/app/service"
	"im-server/config"
)

// Router Gin路由
func Router() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(zaplog.GinLogger())
	r.Use(zaplog.GinRecovery())

	// 初始化业务路由存放的路径
	service.InitRouter(wsUser)
	// 添加不需要登录的路由就可以访问的指令
	notLoginCmd()

	r.GET("/v1/ws", service.WebsocketServer)

	return r
}

// Run 直接运行 Gin Http
func Run() error {
	ipAddr := config.AppIPAddr()
	zaplog.Debugf("▶▶▶启用IP端口:%v", ipAddr)
	return Router().Run(ipAddr)
}
