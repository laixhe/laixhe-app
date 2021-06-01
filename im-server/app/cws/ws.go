package cws

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/laixhe/goutil/zaplog"

	"im-server/servers"
)

// 定义一个 upgrade 类型用于升级 http 为 websocket
var upgrader = websocket.Upgrader{
	HandshakeTimeout: 3 * time.Second,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// WebsocketServer 升级 websocket 服务
func WebsocketServer(c *gin.Context, m *servers.ClientManager) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		zaplog.Errorf("upgrader: %v", err)
		return
	}
	servers.NewClient(conn, m).Start()
}
