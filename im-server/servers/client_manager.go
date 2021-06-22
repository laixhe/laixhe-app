package servers

import (
	"sync"

	"im-server/protoim"
)

// clientManager 客户端连接管理
type clientManager struct {
	clients    *sync.Map    // 保存客户端的链接   client.addr
	users      *sync.Map    // 保存用户登录的链接 client.userId
	userLogin  chan *client // 用户登录处理
	register   chan *client // 注册客户端
	unregister chan *client // 移除注册客户端
	broadcast  chan []byte  // 通知广播消息
	//
	router         *Router              // 业务路由存放的路径
	routerNotLogin map[protoim.CMD]bool // 不需要登录的路由就可以访问的
}

// newClientManager init
func newClientManager() *clientManager {
	return &clientManager{
		clients:    new(sync.Map),
		users:      new(sync.Map),
		userLogin:  make(chan *client, 100),
		register:   make(chan *client, 100),
		unregister: make(chan *client, 100),
		broadcast:  make(chan []byte, 100),
		//
		router:         newRouter(),
		routerNotLogin: make(map[protoim.CMD]bool),
	}
}

// run 连接管理
func (m *clientManager) run() {
	for {
		select {
		case c := <-m.register:
			// 注册客户端的链接
			m.clients.Store(c.addr, c)
		case c := <-m.userLogin:
			// 用户登录处理
			m.users.Store(c.userID, c)
		case c := <-m.unregister:
			// 移除注册客户端的链接
			m.clients.Delete(c.addr)
			if c.isLogin {
				m.users.Delete(c.userID)
			}
		case msg := <-m.broadcast:
			// 广播消息
			// 给每个客户端发消息
			m.clients.Range(func(clientAddr, clientInterface interface{}) bool {
				c, is := clientInterface.(*client)
				if !is {
					// 移除注册客户端的链接
					m.clients.Delete(clientAddr)
					return false
				}
				if c.isClosed {
					// 移除注册客户端的链接
					m.clients.Delete(clientAddr)
					return false
				}
				if c.isLogin {
					_ = c.sendData(msg)
				}
				return true
			})
		}
	}
}

// initRouter 初始化业务路由存放的路径
func (m *clientManager) initRouter(f func(r *Router)) {
	f(m.router)
}

// isNotLoginCmd -
func (m *clientManager) isNotLoginCmd(cmd protoim.CMD) bool {
	return m.routerNotLogin[cmd]
}
