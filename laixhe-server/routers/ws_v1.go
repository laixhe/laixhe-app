package routers

import (
	"github.com/laixhe/laixhe-app/laixhe-server/app/websockets/w_v1"
	"github.com/laixhe/laixhe-app/laixhe-server/servers/ws"
)

// v1 版本
func initWsV1() {

	// login
	ws.RouterSet(1, w_v1.WSLogin)

}
