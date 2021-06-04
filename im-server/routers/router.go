package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/laixhe/goutil/zaplog"

	"im-server/app/cws"
	"im-server/config"
	"im-server/servers"
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

	// 初始化客户端连接管理
	clientManager := servers.NewClientManager()
	go clientManager.Run()
	clientManager.InitRouter(wsUser)

	routerV1 := r.Group("/v1")
	{
		routerV1.GET("/ws", func(c *gin.Context) {
			cws.WebsocketServer(c, clientManager)
		})
	}
	return r
}

// 直接运行 Gin Http
func Run() error {
	ipAddr := config.AppIPAddr()
	zaplog.Debugf("▶▶▶启用IP端口:%v", ipAddr)
	return Router().Run(ipAddr)
}
