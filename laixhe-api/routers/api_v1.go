package routers

import (
	"github.com/laixhe/laixhe-app/laixhe-server/app/controllers"
)

// v1 版本
func initApiV1() {

	rv := "/v1"

	router.GET(rv+"/token", ginLogger(), controllers.Token)

	v := router.Group(rv, ginLogger(), ginJWTAuth())
	{
		v.GET("/ping", controllers.Ping)
	}

}
