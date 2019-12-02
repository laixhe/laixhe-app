package routers

import (
	"github.com/laixhe/laixhe-app/laixhe-server/app/controllers/c_v1"
)

// v1 版本
func initApiV1() {

	rv := "/v1"

	router.GET(rv+"/token", ginLogger(), c_v1.Token)
	router.GET(rv+"/ws", ginLogger(), c_v1.Ws)

	v := router.Group(rv, ginLogger(), ginJWTAuth())
	{
		v.GET("/ping", c_v1.Ping)
	}

}
