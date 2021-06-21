package servers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/laixhe/goutil/zaplog"

	"im-server/protoim"
)

// Servers 服务管理
type Servers struct {
	clientManager *clientManager      // 客户端连接管理
	wsUpgrader    *websocket.Upgrader // 用于升级 http 为 websocket
}

// NewServers init Servers 服务管理
func NewServers() *Servers {
	s := &Servers{
		clientManager: newClientManager(),
		wsUpgrader: &websocket.Upgrader{
			HandshakeTimeout: 3 * time.Second,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	}
	go s.clientManager.run()
	return s
}

// WebsocketUpgrader 定义一个 upgrade 类型用于升级 http 为 websocket
func (s *Servers) WebsocketUpgrader(upgrader *websocket.Upgrader) {
	if upgrader != nil {
		s.wsUpgrader = upgrader
	}
}

// WebsocketServer 由 http 升级 websocket 服务
func (s *Servers) WebsocketServer(c *gin.Context) {
	conn, err := s.wsUpgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		zaplog.Errorf("WebsocketServer upgrader: %v", err)
		return
	}
	newClient(conn, s.clientManager).start()
}

// InitRouter 初始化业务路由存放的路径
func (s *Servers) InitRouter(f func(r *Router)) {
	s.clientManager.initRouter(f)
}

// AddNotLoginCmd 添加不需要登录的路由就可以访问的指令
func (s *Servers) AddNotLoginCmd(cmd protoim.CMD) {
	s.clientManager.routerNotLogin[cmd] = true
}

// IsLogin 用户是否登录
func (s *Servers) IsLogin(userID string) bool {
	_, ok := s.clientManager.clients.Load(userID)
	return ok
}

// SendMessageResponse 发送信息
func (s *Servers) SendMessageResponse(data *protoim.MessageResponse) *protoim.ErrorBase {
	protoBase, err := EnCode(protoim.CMD_C_MESSAGE_RESPONSE, data)
	if err != nil {
		return ErrorMessage(protoim.Error_E_ENCODE, err.Error())
	}
	clientInterface, ok := s.clientManager.users.Load(data.ToId)
	if ok {
		c, is := clientInterface.(*client)
		if !is {
			// 移除注册客户端的链接
			s.clientManager.users.Delete(data.ToId)
			return nil
		}
		if c.isClosed {
			// 移除注册客户端的链接
			s.clientManager.users.Delete(data.ToId)
			return nil
		}
		_ = c.sendData(protoBase)
	}
	return nil
}
