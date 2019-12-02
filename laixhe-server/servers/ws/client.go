package ws

import (
	"encoding/json"
	"errors"
	"log"
	"sync"

	"github.com/gorilla/websocket"

	"github.com/laixhe/laixhe-app/laixhe-server/utils"
)

// 用户连接
type Client struct {
	addr     string          // 客户端地址
	conn     *websocket.Conn // 用户连接
	connTime int64           // 连接事件
	send     chan []byte     // 待发送的数据
	isClosed bool            // 当前连接的关闭状态
	lock     sync.Mutex      // 锁-关闭状态
}

// 初始化用户连接
func NewClient(addr string, conn *websocket.Conn, connTime int64) *Client {

	return &Client{
		addr:     addr,
		conn:     conn,
		connTime: connTime,
		send:     make(chan []byte, 100),
	}

}

// 开始处理连接的读写
func (this *Client) Start() {

	// 处理当前客户端发送信息
	go this.writeClientChan()

	// 处理接收当前客户端发送过来的数据
	go this.readClientChan()

}

// 处理接收用户发送过来的数据
func (this *Client) readClientChan() {
	defer this.Stop()

	for {

		// 读取 ws 中的数据
		_, message, err := this.conn.ReadMessage()
		if err != nil {
			utils.ZapSugar().Infow("Websocket Read Message: " + err.Error())
			break
		}

		err = this.processData(message)
		if err != nil {
			utils.ZapSugar().Infow("Websocket Read Message ProcessData: " + err.Error())
			break
		}

	}

}

// 专门给当前客户端发送信息
func (this *Client) writeClientChan() {
	defer this.Stop()

	for msg := range this.send {

		err := this.conn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			utils.ZapSugar().Infow("Websocket Write Message: " + err.Error())
			break
		}

	}

}

// 给当前客户端发送信息
func (this *Client) SendClient(data []byte) error {

	this.lock.Lock()
	defer this.lock.Unlock()

	// 判断当前链接是否已经关闭
	if this.isClosed {
		return errors.New("client close")
	}

	this.send <- data

	return nil
}

// 停止连接
func (this *Client) Stop() {

	this.lock.Lock()
	defer this.lock.Unlock()

	// 判断当前链接是否已经关闭
	if this.isClosed {
		return
	}
	this.isClosed = true

	// 关闭连接
	this.conn.Close()

	// 关闭 - 用户发送数据的管道
	close(this.send)

}

// 处理数据
func (this *Client) processData(message []byte) error {

	log.Println("读取客户端数据 处理:", string(message))

	cmd := &WSCmd{}
	err := json.Unmarshal(message, cmd)
	if err != nil {
		return err
	}

	log.Println("读取客户端数据 cmd=", cmd)

	// 分配路由
	err = RouterGet(NewRequest(this, message, cmd.Cmd))
	if err != nil {
		return err
	}

	return nil
}
