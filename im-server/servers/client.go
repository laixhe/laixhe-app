package servers

import (
	"fmt"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/laixhe/goutil/zaplog"

	"im-server/config"
	"im-server/protoim"
	"im-server/utils"
)

// ReadTime 读超时(秒)，同时也是心跳超时
var ReadTime = time.Second * config.WebsocketReadTime()

// WriteTime 写超时(秒)
var WriteTime = time.Second * config.WebsocketWriteTime()

// Client 客户端连接
type Client struct {
	conn       *websocket.Conn // 客户端连接
	addr       string          // 客户端地址
	send       chan []byte     // 待发送的数据
	connTime   time.Time       // 连接时间
	isClosed   bool            // 当前连接的关闭状态
	lockClosed *sync.Mutex     // 当前连接的关闭状态-锁
	//
	appOsId   protoim.AppOs // 用户登录的平台 Id
	userId    string        // 用户 Id
	userKey   string        // 用户平台连接 Key
	loginTime time.Time     // 用户登录时间
	isLogin   bool          // 用户是否登录
	//
	manager *ClientManager // 客户端连接管理
}

// NewClient init
func NewClient(conn *websocket.Conn, manager *ClientManager) *Client {
	return &Client{
		conn:       conn,
		addr:       conn.RemoteAddr().String(),
		send:       make(chan []byte, 100),
		connTime:   time.Now(),
		isClosed:   false,
		lockClosed: new(sync.Mutex),
		//
		manager: manager,
	}
}

// Start 开始处理连接的读写
func (c *Client) Start() {
	// 对 ping、pong 侦的处理
	// c.conn.SetPongHandler(func(string) error {
	// 	_ = c.conn.SetReadDeadline(time.Now().Add(ReadTime))
	// 	return nil
	// })
	// c.conn.SetPingHandler(func(string) error {
	// 	_ = c.conn.SetReadDeadline(time.Now().Add(ReadTime))
	// 	return nil
	// })
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
		// 读取 ws 中的数据
		messageType, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				zaplog.Errorf("addr read11=%v err=%v", c.addr, err)
				break
			}
			zaplog.Errorf("addr read22=%v err=%v", c.addr, err)
			break
		}
		messageBase, err := DeCodeMessageBase(message)
		if err != nil {
			zaplog.Errorf("addr read=%v err=%v type=%v", c.addr, err, messageType)
			if d, e := ErrorMessageDeCode(); e == nil {
				c.send <- d
			} else {
				zaplog.Errorf("addr read=%v err=%v type=%v", c.addr, e, messageType)
			}
			break
		}
		// 路由业务
		errMsg := c.manager.router.Get(&Context{
			conn: c,
			data: messageBase.Data,
			cmd:  messageBase.Cmd,
		})
		if errMsg != nil {
			zaplog.Errorf("addr read=%v err=%v type=%v", c.addr, errMsg, messageType)
			if d, e := EnErrorMessage(errMsg); e == nil {
				c.send <- d
			} else {
				zaplog.Errorf("addr read=%v err=%v type=%v", c.addr, e, messageType)
			}
			continue
		}
		// 设置 读超时(秒)，同时也是心跳超时
		_ = c.conn.SetReadDeadline(time.Now().Add(ReadTime))
	}
}

// writeClientChan 专门给当前客户端发送信息
func (c *Client) writeClientChan() {
	defer c.Stop()
	for msg := range c.send {
		// 设置 写超时(秒)
		_ = c.conn.SetWriteDeadline(time.Now().Add(WriteTime))
		err := c.conn.WriteMessage(websocket.BinaryMessage, msg)
		if err != nil {
			zaplog.Errorf("addr write=%v err=%v", c.addr, err)
			return
		}
		fmt.Printf("T--------- write %v | %v\n", msg, len(msg))
	}
}

// userLogin 用户登录
func (c *Client) userLogin(appOsId protoim.AppOs, userId string) {
	if c.isClosed {
		return
	}
	c.appOsId = appOsId        // 用户登录的平台 Id
	c.userId = userId          // 用户 Id
	c.userKey = c.getUserKey() // 用户平台连接 Key
	c.loginTime = time.Now()   // 用户登录时间
	c.isLogin = true           // 用户是否登录
	// 用户登录处理
	c.manager.userLogin <- c
}

// getUserConnKey 用户平台连接key
func (c *Client) getUserKey() string {
	if c.userKey == "" {
		c.userKey = utils.GetUserKey(c.appOsId, c.userId)
	}
	return c.userKey
}
