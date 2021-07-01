package routers

import (
	"im-server/app/cws"
	"im-server/protoim"
	"im-server/servers"
)

// wsUser 初始化业务路由
func wsUser(r *servers.Router) {
	r.Set(protoim.CMD_C_USER_LOGIN_REQUEST, cws.UserLoginRequest)
	r.Set(protoim.CMD_C_GET_USER_REQUEST, cws.GetUserRequest)
	r.Set(protoim.CMD_C_GET_FRIENDS_REQUEST, cws.GetFriendsRequest)
	r.Set(protoim.CMD_C_MESSAGE_SEND_REQUEST, cws.MessageSendRequest)
}
