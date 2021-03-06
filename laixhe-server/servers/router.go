package servers

import (
	"fmt"

	"github.com/laixhe/laixhe-app/laixhe-server/protoapi"
)

// 业务路由存放的路径
type Router struct {
	path map[protoapi.CMD]func(*Request)
}

var router = &Router{path: make(map[protoapi.CMD]func(*Request))}

// 设置业务路由
func RouterSet(cmd protoapi.CMD, fun func(*Request)) {

	if fun == nil {
		panic("router can not be nil")
	}

	router.path[cmd] = fun
}

// 获取业务路由
func RouterGet(req *Request) error {

	fun, ok := router.path[req.cmd]
	if !ok {
		return fmt.Errorf("cmd no find: %d", req.cmd)
	}

	go fun(req)

	return nil
}
