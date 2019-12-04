package c_v1

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"

	"github.com/laixhe/laixhe-app/laixhe-server/servers/ws"
	"github.com/laixhe/laixhe-app/laixhe-server/utils"
	"github.com/laixhe/laixhe-app/laixhe-server/utils/logs"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:   1024,
	WriteBufferSize:  1024,
	HandshakeTimeout: 5 * time.Second,
	CheckOrigin: func(r *http.Request) bool {
		return true // 取消 ws 跨域校验
	},
}

// websocket
func Ws(c *gin.Context) {

	// 升级 get 请求为 webSocket 协议
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {

		utils.GinJsonResponseMsg(c, utils.ERROR_WEBSOCKET, err.Error())
		logs.Debug("ERROR_WEBSOCKET: ", err)

		return
	}

	// 初始化用户连接
	client := ws.NewClient(conn.RemoteAddr().String(), conn, time.Now().Unix())

	// 开始处理连接的读写
	client.Start()

}
