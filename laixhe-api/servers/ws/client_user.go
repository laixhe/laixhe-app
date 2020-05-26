package ws

import (
	"fmt"
	"time"
)

const (
	HeartbeatExpirationTime = 120 // 用户连接超时时间
)

type ClientUser struct {
	client        *Client // 用户连接
	AppId         AppOS   // 登录的平台Id app/web/ios
	UserId        string  // 用户Id，用户登录以后才有
	HeartbeatTime int64   // 用户上次心跳时间
	LoginTime     int64   // 登录时间
}

func NewClientUser(client *Client, appId AppOS, userId string) *ClientUser {
	return &ClientUser{
		client:        client,
		AppId:         appId,
		UserId:        userId,
		HeartbeatTime: time.Now().Unix(),
		LoginTime:     time.Now().Unix(),
	}
}

// 获取用户 key
func (this *ClientUser) GetUserKey() string {
	return GetUserKey(this.AppId, this.UserId)
}

// 获取用户 key
func GetUserKey(appId AppOS, userId string) string {
	return fmt.Sprintf("%d_%s", appId, userId)
}
