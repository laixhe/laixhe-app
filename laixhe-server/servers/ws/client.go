package ws

import (
	"encoding/json"
	"errors"
	"log"
	"net"
	"sync"
	"time"

	"github.com/gorilla/websocket"

	"github.com/laixhe/laixhe-app/laixhe-server/utils/logs"
)

// 用户连接
type Client struct {
	addr       string          // 客户端地址
	conn       *websocket.Conn // 用户连接
	connTime   int64           // 连接时间
	send       chan []byte     // 待发送的数据
	isClosed   bool            // 当前连接的关闭状态
	lockClosed sync.Mutex      // 锁-关闭状态
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

		//
		this.conn.SetReadDeadline(time.Now().Add(time.Second * 2))

		// 读取 ws 中的数据
		_, message, err := this.conn.ReadMessage()
		if err != nil {

			// 判断是不是超时
			if netErr, ok := err.(net.Error); ok {
				if netErr.Timeout() {
					logs.Errorf("Websocket ReadMessage timeout remote: %v\n", this.conn.RemoteAddr())
					break
				}
			}

			// 其他错误，如果是 1001 和 1000 就不打印日志
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseNormalClosure) {
				logs.Errorf("Websocket ReadMessage other remote:%v error: %v \n", this.conn.RemoteAddr(), err)
			}

			break
		}

		err = this.processData(message)
		if err != nil {
			logs.Error("Websocket Read Message ProcessData: ", err)
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
			logs.Error("Websocket Write Message: ", err)
			break
		}

	}

}

// 给当前客户端发送信息
func (this *Client) SendClient(data []byte) error {

	this.lockClosed.Lock()
	defer this.lockClosed.Unlock()

	// 判断当前链接是否已经关闭
	if this.isClosed {
		return errors.New("client close")
	}

	this.send <- data

	return nil
}

// 停止连接
func (this *Client) Stop() {

	this.lockClosed.Lock()
	defer this.lockClosed.Unlock()

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
	return nil
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
