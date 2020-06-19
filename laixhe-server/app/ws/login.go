package ws

import (
	"fmt"
	"github.com/laixhe/laixhe-app/laixhe-server/protoapi"
	"github.com/laixhe/laixhe-app/laixhe-server/servers"
	"strconv"
	"time"
)

// 登录请求
//func (this *protoapi.LoginRequest) GetCmd() uint {
//	return 3
//}
//
////====================================================
//
//// 登录响应数据
//func (this *protoapi.LoginResponse) GetCmd() uint {
//	return 3
//}

//====================================================

func WSLogin(w *servers.Request) {

	login := new(protoapi.LoginRequest)
	err := w.Message(login)
	if err != nil {
		fmt.Println("T------- GetMsg:", err)
		return
	}

	fmt.Printf("T---. login: %#v\n", login)

	err = w.Send(&protoapi.LoginResponse{Message: "--user_id--" + strconv.Itoa(int(time.Now().Unix()))})
	if err != nil {
		fmt.Println("T------- Send:", err)
	}

}
