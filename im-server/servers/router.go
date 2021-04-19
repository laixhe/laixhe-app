package servers

import (
	"fmt"
	"im-server/protoim"
)

// Router 业务路由存放的路径
type Router struct {
	path map[protoim.CMD]func(*Context)
}

// NewRouter init
func NewRouter() *Router {
	return &Router{path: make(map[protoim.CMD]func(*Context))}
}

// Set 设置业务路由
func (r *Router) Set(cmd protoim.CMD, fun func(*Context)) {
	if fun == nil {
		panic("router can not be nil")
	}
	r.path[cmd] = fun
}

// Get 获取业务路由
func (r *Router) Get(req *Context) error {
	fun, ok := r.path[req.cmd]
	if !ok {
		return fmt.Errorf("cmd no find: %d", req.cmd)
	}
	go fun(req)
	return nil
}
