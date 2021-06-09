package service

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"

	"im-server/protoim"
	"im-server/servers"
)

// wsServer websocket 服务管理
var wsServer = servers.NewServers()

// WebsocketUpgrader 定义一个 upgrade 类型用于升级 http 为 websocket
func WebsocketUpgrader(upgrader *websocket.Upgrader) {
	wsServer.WebsocketUpgrader(upgrader)
}

// WebsocketServer 由 http 升级 websocket 服务
func WebsocketServer(c *gin.Context) {
	wsServer.WebsocketServer(c)
}

// InitRouter 初始化业务路由存放的路径
func InitRouter(f func(r *servers.Router)) {
	wsServer.InitRouter(f)
}

// AddNotLoginCmd 添加不需要登录的路由就可以访问的指令
func AddNotLoginCmd(cmd protoim.CMD) {
	wsServer.AddNotLoginCmd(cmd)
}

// 用户是否登录
func IsLogin(userID string) bool {
	return wsServer.IsLogin(userID)
}

// 发送信息
func SendMessageResponse(data *protoim.MessageResponse) *protoim.ErrorInfo {
	return wsServer.SendMessageResponse(data)
}
