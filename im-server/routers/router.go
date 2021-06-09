package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/laixhe/goutil/zaplog"

	"im-server/app/service"
	"im-server/config"
)

// Router Gin路由
func Router() *gin.Engine {
	if !config.IsConsoleLog() {
		// 禁用控制台输出
		gin.DisableConsoleColor()
	}

	if config.IsAppDebug() {
		gin.SetMode(gin.DebugMode)
	}
	r := gin.New()
	// 中间件
	// if config.IsConsoleLog() {
	// 	r.Use(gin.Logger())
	// 	r.Use(gin.Recovery())
	// } else {
	// 	r.Use(zaplog.GinLogger())
	// 	r.Use(zaplog.GinRecovery())
	// }
	r.Use(zaplog.GinLogger())
	r.Use(zaplog.GinRecovery())

	// 初始化业务路由存放的路径
	service.InitRouter(wsUser)
	// 添加不需要登录的路由就可以访问的指令
	notLoginCmd()

	routerV1 := r.Group("/v1")
	{
		routerV1.GET("/ws", service.WebsocketServer)
	}
	return r
}

// 直接运行 Gin Http
func Run() error {
	ipAddr := config.AppIPAddr()
	zaplog.Debugf("▶▶▶启用IP端口:%v", ipAddr)
	return Router().Run(ipAddr)
}
