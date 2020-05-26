package ws

import (
	"errors"
	"log"
	"sync"
)

var UserManager = NewClientManager()

// 用户连接管理
type ClientManager struct {
	clients   *sync.Map   // 保存用户的链接
	delClient chan string // 通知要删除用户的下标
	register  chan *Client
	destroy   chan *Client
	broadcast chan []byte // 通知广播消息
}

func NewClientManager() *ClientManager {

	manager := &ClientManager{
		clients:   new(sync.Map),
		delClient: make(chan string, 100),
		broadcast: make(chan []byte, 100),
	}

	go manager.deleteClientChan()
	go manager.broadcastChan()

	return manager
}

// 处理删除用户
func (this *ClientManager) deleteClientChan() {

	for connID := range this.delClient {
		this.clients.Delete(connID)
		log.Println(connID, ": 链接退出")
	}

}

// 广播消息
func (this *ClientManager) broadcastChan() {

	for msg := range this.broadcast {

		// 给每个用户发消息
		this.clients.Range(func(connID, clientUser interface{}) bool {

			user, is := clientUser.(*ClientUser)
			if !is {
				this.DeleteClient(connID.(string))
				return false
			}

			err := user.client.SendClient(msg)
			if err != nil {
				this.DeleteClient(connID.(string))
				return false
			}

			return true

		})

	}

}

// 处理删除用户
func (this *ClientManager) DeleteClient(connID string) {
	this.delClient <- connID
}

// 广播消息
func (this *ClientManager) BroadcastClient(data WsDataInterface) error {

	// 数据打包序列化
	wsData := NewWsData(data)
	code, err := wsData.Encode()
	if err != nil {
		return err
	}

	this.broadcast <- code

	return nil
}

// 保存在线用户
func (this *ClientManager) SaveUser(client *ClientUser){
	this.clients.Store(client.GetUserKey(), client)
}

// 客户端发送消息
func (this *ClientManager) Send(connID string, data WsDataInterface) error {

	// 数据打包序列化
	wsData := NewWsData(data)
	code, err := wsData.Encode()
	if err != nil {
		return err
	}

	// 获取客户端
	clientUser,ok := this.clients.Load(connID)
	if ok {

		user, is := clientUser.(*ClientUser)
		if !is {
			this.DeleteClient(connID)
			return errors.New("client load 'ClientUser' Failure of type conversion")
		}

		err := user.client.SendClient(code)
		if err != nil {
			this.DeleteClient(connID)
			return err
		}

		return nil
	}

	return errors.New("client not load")
}

// 是否存在用户
func (this *ClientManager) IsUser(connID string) bool {
	// 获取客户端
	_,ok := this.clients.Load(connID)
	return ok
}
