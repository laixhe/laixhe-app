package servers

import (
	"fmt"

	"im-server/protoim"
)

// Router 业务路由存放的路径
type Router struct {
	path map[protoim.CMD]func(*Context)
}

// newRouter init
func newRouter() *Router {
	r := &Router{path: make(map[protoim.CMD]func(*Context))}

	// 添加心跳路由
	r.Set(protoim.CMD_C_PING, ping)
	return r
}

// Set 设置业务路由
func (r *Router) Set(cmd protoim.CMD, f func(*Context)) {
	if f == nil {
		panic("router can not be nil")
	}
	r.path[cmd] = f
}

// Get 获取业务路由
func (r *Router) Get(req *Context) *protoim.ErrorInfo {
	fun, ok := r.path[req.cmd]
	if !ok {
		return ErrorMessage(protoim.Error_E_ROUTE_NOT_EXIST, fmt.Sprintf("cmd no find: %d", req.cmd))
	}
	go fun(req)
	return nil
}
