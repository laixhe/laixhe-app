package servers

import (
	"fmt"
	"im-server/protoim"
	"sync"
	"time"

	"google.golang.org/protobuf/proto"

	"github.com/gorilla/websocket"
	"github.com/laixhe/goutil/zaplog"

	"im-server/config"
)

// ReadTime 读超时(秒)，同时也是心跳超时
var ReadTime = time.Second * config.WebsocketReadTime()

// WriteTime 写超时(秒)
var WriteTime = time.Second * config.WebsocketWriteTime()

// Client 客户端连接
type Client struct {
	conn       *websocket.Conn // 客户端连接
	addr       string          // 客户端地址
	connTime   time.Time       // 连接时间
	send       chan []byte     // 待发送的数据
	isClosed   bool            // 当前连接的关闭状态
	lockClosed *sync.Mutex     // 锁-关闭状态
	manager    *ClientManager  // 客户端连接管理
	router     *Router         // 业务路由存放的路径
}

// NewClient init
func NewClient(conn *websocket.Conn, manager *ClientManager, router *Router) *Client {
	return &Client{
		addr:       conn.RemoteAddr().String(),
		conn:       conn,
		connTime:   time.Now(),
		send:       make(chan []byte, 100),
		lockClosed: new(sync.Mutex),
		manager:    manager,
		router:     router,
	}
}

// Start 开始处理连接的读写
func (c *Client) Start() {
	// 对 ping、pong 侦的处理
	c.conn.SetPongHandler(func(string) error {
		_ = c.conn.SetReadDeadline(time.Now().Add(ReadTime))
		return nil
	})
	c.conn.SetPingHandler(func(string) error {
		_ = c.conn.SetReadDeadline(time.Now().Add(ReadTime))
		return nil
	})
	// 注册客户端的链接
	c.manager.register <- c
	go c.writeClientChan()
	go c.readClientChan()
}

// Stop 停止连接
func (c *Client) Stop() {
	c.lockClosed.Lock()
	defer c.lockClosed.Unlock()
	// 判断当前链接是否已经关闭
	if c.isClosed {
		return
	}
	c.isClosed = true
	// 移除注册客户端的链接
	c.manager.unregister <- c
	// 关闭-客户端发送数据的管道
	close(c.send)
	// 关闭连接
	_ = c.conn.Close()
}

// readClientChan 处理接收客户端发送过来的数据
func (c *Client) readClientChan() {
	defer c.Stop()
	for {
		// 设置 读超时(秒)，同时也是心跳超时
		_ = c.conn.SetReadDeadline(time.Now().Add(ReadTime))
		// 读取 ws 中的数据
		messageType, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				zaplog.Errorf("addr read: %v : %v", c.addr, err)
			}
			break
		}

		fmt.Println("T--------- read:", messageType)
		messageBase := new(protoim.MessageBase)
		err = proto.Unmarshal(message, messageBase)
		if err != nil {
			fmt.Println("T--------- read messageBase err: ", err)
			return
		}
		fmt.Println("T--------- messageBase read:", messageBase)
		getUserInfoRequest := new(protoim.GetUserInfoRequest)
		err = proto.Unmarshal(messageBase.Data, getUserInfoRequest)
		if err != nil {
			fmt.Println("T--------- read getUserInfoRequest err: ", err)
			return
		}
		fmt.Println("T--------- getUserInfoRequest read:", getUserInfoRequest)
	}
}

// writeClientChan 专门给当前客户端发送信息
func (c *Client) writeClientChan() {
	defer c.Stop()
	for msg := range c.send {
		// 设置 写超时(秒)
		_ = c.conn.SetWriteDeadline(time.Now().Add(WriteTime))
		err := c.conn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			zaplog.Errorf("addr write: %v : %v", c.addr, err)
			return
		}
		fmt.Println("T--------- write", string(msg))
	}
}
