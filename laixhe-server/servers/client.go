package servers

import (
	"encoding/binary"
	"errors"
	"github.com/laixhe/laixhe-app/laixhe-server/protoapi"
	"net"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/laixhe/goutil/zap_log"
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

		// 设置 读超时
		this.conn.SetReadDeadline(time.Now().Add(time.Second * 2))

		// 读取 ws 中的数据
		_, message, err := this.conn.ReadMessage()
		if err != nil {

			// 判断是不是超时
			if netErr, ok := err.(net.Error); ok {
				if netErr.Timeout() {
					zap_log.Errorf("addr: %v : %v : %v", this.conn.RemoteAddr() , protoapi.ErrorCode_E_WEBSOCKET_TIMEOUT, err)
					break
				}
			}

			// 其他错误，如果是 1001 和 1000 就不打印日志
			//if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseNormalClosure) {
			//	zap_log.Errorf("Websocket ReadMessage other remote:%v error: %v \n", this.conn.RemoteAddr(), err)
			//}

			zap_log.Errorf("addr: %v : %v : %v", this.conn.RemoteAddr() , protoapi.ErrorCode_E_WEBSOCKET_READ, err)
			break
		}

		err = this.processData(message)
		if err != nil {
			zap_log.Errorf("addr: %v : %v : %v", this.conn.RemoteAddr() , protoapi.ErrorCode_E_DECODE, err)
			break
		}

	}

}

// 专门给当前客户端发送信息
func (this *Client) writeClientChan() {
	defer this.Stop()

	for msg := range this.send {

		err := this.conn.WriteMessage(websocket.BinaryMessage, msg)
		if err != nil {
			zap_log.Errorf("addr: %v : %v : %v", this.conn.RemoteAddr() , protoapi.ErrorCode_E_WEBSOCKET_WRITE, err)
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

	cmd := binary.BigEndian.Uint32(message)
	zap_log.Debugf("addr: %v : get cmd : %v", this.conn.RemoteAddr(), cmd)

	// 分配路由
	err := RouterGet(NewRequest(this, message[4:], protoapi.CMD(cmd)))
	if err != nil {
		zap_log.Errorf("addr: %v : %v : %v", this.conn.RemoteAddr(), protoapi.ErrorCode_E_ROUTING_NOT_EXIST, err)
		return err
	}

	return nil
}
