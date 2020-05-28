package ws

import (
	"fmt"
	"strconv"
	"time"

	"github.com/laixhe/laixhe-app/laixhe-server/servers"
)

// 请求登录数据
type LoginRequest struct {
	Token string `json:"token"`  // 验证用户是否登录
	AppId uint   `json:"app_id"` // 登录的平台 AppOS
}

func (this *LoginRequest) GetCmd() uint {
	return 3
}

//====================================================

// 响应登录数据
type LoginResponse struct {
	UserId string `json:"user_id"` // 用户id
}

func (this *LoginResponse) GetCmd() uint {
	return 3
}

//====================================================

func WSLogin(w *servers.Request) {

	login := &LoginRequest{}
	l, err := w.Message(login)
	if err != nil {
		fmt.Println("T------- GetMsg:", err)
	}

	fmt.Printf("T---. login l: %#v\n", l)
	fmt.Printf("T---. login: %#v\n", login)

	err = w.Send(&LoginResponse{UserId: "--user_id--" + strconv.Itoa(int(time.Now().Unix()))})
	if err != nil {
		fmt.Println("T------- Send:", err)
	}

}
