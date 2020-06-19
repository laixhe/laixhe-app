package routers

import (
	"github.com/laixhe/laixhe-app/laixhe-server/app/ws"
	"github.com/laixhe/laixhe-app/laixhe-server/protoapi"
	"github.com/laixhe/laixhe-app/laixhe-server/servers"
)

// v1 版本
func initWsV1() {

	// login
	servers.RouterSet(protoapi.CMD_C_LOGIN, ws.WSLogin)

}
