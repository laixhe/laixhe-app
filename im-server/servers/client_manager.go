package servers

import (
	"sync"
)

// ClientManager 客户端连接管理
type ClientManager struct {
	clients    *sync.Map    // 保存客户端的链接   client.addr
	users      *sync.Map    // 保存用户登录的链接 client.userConnKey
	userLogin  chan *Client // 用户登录处理
	register   chan *Client // 注册客户端
	unregister chan *Client // 移除注册客户端
	broadcast  chan []byte  // 通知广播消息
}

// NewClientManager 初始化客户端连接管理
func NewClientManager() *ClientManager {
	return &ClientManager{
		clients:    new(sync.Map),
		users:      new(sync.Map),
		userLogin:  make(chan *Client, 100),
		register:   make(chan *Client, 100),
		unregister: make(chan *Client, 100),
		broadcast:  make(chan []byte, 100),
	}
}

// Run 连接管理
func (m *ClientManager) Run() {
	for {
		select {
		case client := <-m.register:
			// 注册客户端的链接
			m.clients.Store(client.addr, client)
		case client := <-m.userLogin:
			// 用户登录处理
			m.users.Store(client.userKey, client)
		case client := <-m.unregister:
			// 移除注册客户端的链接
			m.clients.Delete(client.addr)
			if client.isLogin {
				m.users.Delete(client.userKey)
			}
		case message := <-m.broadcast:
			// 广播消息
			// 给每个客户端发消息
			m.clients.Range(func(clientAddr, clientInterface interface{}) bool {
				client, is := clientInterface.(*Client)
				if !is {
					// 移除注册客户端的链接
					m.clients.Delete(clientAddr)
					return false
				}
				if client.isClosed {
					// 移除注册客户端的链接
					m.clients.Delete(clientAddr)
					return false
				}
				client.send <- message
				return true
			})
		}
	}
}
