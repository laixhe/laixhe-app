package routers

import (
	"bytes"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/laixhe/goutil"

	"github.com/laixhe/laixhe-app/laixhe-server/config"
	"github.com/laixhe/laixhe-app/laixhe-server/utils"
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

	router.Static("/static", "./static")

	// API V1
	initApiV1()

	// WS V1
	initWsV1()

	return router
}

// gin 日志中间件
func ginLogger() gin.HandlerFunc {

	return func(c *gin.Context) {

		urlString := c.Request.URL.String()
		method := c.Request.Method
		requestId := c.Request.Header.Get("x-request-id")
		if requestId == "" {
			requestId = goutil.UUidNew()
		}

		c.Header("x-request-id", requestId)
		c.Set("x-request-id", requestId)

		if method == "GET" || method == "HEAD" {

			utils.ZapSugar().Infow("gin_log", utils.REQUEST_ID, requestId, "url", urlString)
			return
		}

		data, err := c.GetRawData()
		if err != nil {

			c.Abort()
			return
		}

		if len(data) > 0 {

			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))
			utils.ZapSugar().Infow("gin_log", utils.REQUEST_ID, requestId, "url", urlString, "body", string(data))
		}

		c.Next()

	}
}

// gin Token 验证
func ginJWTAuth() gin.HandlerFunc {

	return func(c *gin.Context) {

		aut := c.Request.Header.Get("Authorization")
		if aut == "" {

			utils.GinJsonResponseCode(c, utils.ERROR_AUTH)

			c.Abort()
			return
		}

		token, err := utils.TokenCheck(aut, config.GetApp().Token)
		if err != nil {

			utils.GinJsonResponseCode(c, utils.ERROR_AUTH)

			c.Abort()
			return
		}

		c.Set("token", token)

		c.Next()
	}

}

// 直接运行 gin http
func Run(addr, mode string) error {
	InitRouter(mode)
	return router.Run(addr)
}
