package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/laixhe/laixhe-app/laixhe-server/app/ws"

	"github.com/laixhe/laixhe-app/laixhe-server/config"
)

var router *gin.Engine

// 对外路由可扩展
func InitRouter(mode string) *gin.Engine {

	// 设置运行级别
	switch mode {
	case "debug":
		mode = gin.DebugMode
	case "test":
		mode = gin.TestMode
	case "production":
		mode = gin.ReleaseMode
	default:
		mode = gin.DebugMode
	}

	gin.SetMode(mode)

	// 初始化GIN
	router = gin.Default()

	router.GET("/v1/ws", ws.Ws)

	// WS V1
	initWsV1()

	return router
}

// 直接运行 gin http
func Run() error {
	InitRouter(config.GetApp().AppMode)
	return router.Run(config.GetHttp().Addr)
}
