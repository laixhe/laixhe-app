package routers

import (
	"github.com/laixhe/laixhe-app/laixhe-server/app/ws"
	"github.com/laixhe/laixhe-app/laixhe-server/servers"
)

// v1 版本
func initWsV1() {

	// login
	servers.RouterSet(1, ws.WSLogin)

}
