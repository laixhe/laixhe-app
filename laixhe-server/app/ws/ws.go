package ws

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/laixhe/goutil/zap_log"

	"github.com/laixhe/laixhe-app/laixhe-server/servers"
	"github.com/laixhe/laixhe-app/laixhe-server/utils"
)

// 定义一个 upgrade 类型用于升级 http 为 websocket
var upgrader = websocket.Upgrader{
	ReadBufferSize:   1024,            // 指定 io 操作的缓存大小，如果不指定就会自动分配
	WriteBufferSize:  1024,            // 写数据操作的缓存池，如果没有设置值，write buffers 将会分配到链接生命周期里。
	HandshakeTimeout: 5 * time.Second, // 指定升级 websocket 握手完成的超时时间
	CheckOrigin: func(r *http.Request) bool {

		// 请求检查函数，用于统一的链接检查，以防止跨站点请求伪造。如果不检查，就设置一个返回值为true的函数。
		if r.Method != "GET" {
			zap_log.Error("method is not GET")
			//return false
		}

		// 取消 ws 跨域校验
		return true
	},
}

// get websocket
func Ws(c *gin.Context) {

	// 升级 http get 请求为 webSocket 协议
	// 如果升级失败，则使用 HTTP 错误响应回复客户端
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {

		utils.GinJsonResponseMsg(c, utils.ERROR_WEBSOCKET, err.Error())
		zap_log.Info("ERROR_WEBSOCKET" + err.Error())

		return
	}

	// 初始化用户连接
	client := servers.NewClient(conn.RemoteAddr().String(), conn, time.Now().Unix())

	// 开始处理连接的读写
	client.Start()

}
