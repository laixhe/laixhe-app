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

// WsReadTime 读超时(秒)，同时也是心跳超时
var WsReadTime = time.Second * config.WebsocketReadTime()

// WsWriteTime 写超时(秒)
var WsWriteTime = time.Second * config.WebsocketWriteTime()

// client 客户端连接
type client struct {
	conn       *websocket.Conn // 客户端连接
	addr       string          // 客户端地址
	send       chan []byte     // 待发送的数据
	connTime   time.Time       // 连接时间
	isClosed   bool            // 当前连接的关闭状态
	lockClosed *sync.Mutex     // 当前连接的关闭状态-锁
	//
	osID      protoim.AppOs // 用户登录的平台 Id
	userID    string        // 用户 Id
	userKey   string        // 用户平台连接 Key
	loginTime time.Time     // 用户登录时间
	isLogin   bool          // 用户是否登录
	//
	manager *clientManager // 客户端连接管理
}

// newClient init
func newClient(conn *websocket.Conn, manager *clientManager) *client {
	return &client{
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

// start 开始处理连接的读写
func (c *client) start() {
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

// stop 停止连接
func (c *client) stop() {
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

// sendData 发送数据
func (c *client) sendData(data []byte) bool {
	c.lockClosed.Lock()
	defer c.lockClosed.Unlock()
	// 判断当前链接是否已经关闭
	if c.isClosed {
		return false
	}
	c.send <- data
	return true
}

// readClientChan 处理接收客户端发送过来的数据
func (c *client) readClientChan() {
	defer c.stop()
	for {
		// 读取 ws 中的数据
		messageType, message, err := c.conn.ReadMessage()
		if err != nil {
			// if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
			// 	zaplog.Errorf("addr read11=%v err=%v", c.addr, err)
			// 	break
			// }
			zaplog.Errorf("addr read=%v err=%v", c.addr, err)
			break
		}
		messageBase, e := DeCodeMessageBase(message)
		if e != nil {
			zaplog.Errorf("addr read=%v code=%v err=%v type=%v", c.addr, e.Code, e.Msg, messageType)
			break
		}
		if !c.isLogin {
			if !c.manager.isNotLoginCmd(messageBase.Cmd) {
				c.sendErrorInfo(ErrorMessage(protoim.E_NOT_LOGIN))
				continue
			}
		}
		// 路由业务
		e = c.manager.router.Get(&Context{
			conn: c,
			data: messageBase.Data,
			cmd:  messageBase.Cmd,
		})
		if e != nil {
			c.sendErrorInfo(e)
			zaplog.Errorf("addr read=%v code=%v err=%v type=%v", c.addr, e.Code, e.Msg, messageType)
			continue
		}
		// 设置 读超时(秒)，同时也是心跳超时
		_ = c.conn.SetReadDeadline(time.Now().Add(WsReadTime))
	}
}

// writeClientChan 专门给当前客户端发送信息
func (c *client) writeClientChan() {
	defer c.stop()
	for msg := range c.send {
		// 设置 写超时(秒)
		_ = c.conn.SetWriteDeadline(time.Now().Add(WsWriteTime))
		err := c.conn.WriteMessage(websocket.BinaryMessage, msg)
		if err != nil {
			zaplog.Errorf("addr write=%v err=%v", c.addr, err)
			return
		}
		fmt.Printf("T--------- write %v | %v\n", msg, len(msg))
	}
}

// userLogin 用户登录
func (c *client) userLogin(osID protoim.AppOs, userId string) {
	if c.isClosed {
		return
	}
	c.osID = osID              // 用户登录的平台 Id
	c.userID = userId          // 用户 Id
	c.userKey = c.getUserKey() // 用户平台连接 Key
	c.loginTime = time.Now()   // 用户登录时间
	c.isLogin = true           // 用户是否登录
	// 用户登录处理
	c.manager.userLogin <- c
}

// getUserConnKey 用户平台连接key
func (c *client) getUserKey() string {
	if c.userKey == "" {
		c.userKey = utils.GetUserKey(c.osID, c.userID)
	}
	return c.userKey
}

// sendErrorInfo 发送错误信息
func (c *client) sendErrorInfo(e *protoim.ErrorInfo) {
	if d, err := EnCodeErrorMessage(e); err == nil {
		_ = c.sendData(d)
	}
}
