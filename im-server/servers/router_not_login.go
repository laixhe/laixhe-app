package servers

import (
	"im-server/protoim"
)

type RouterNotLogin map[protoim.CMD]bool

// NewRouterNotLogin 不需要登录的路由
func NewRouterNotLogin() RouterNotLogin {
	return RouterNotLogin{
		// protoim.CMD_C_ERROR: true,
	}
}
