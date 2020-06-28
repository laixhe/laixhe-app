package ws

import (
	"fmt"
	"github.com/laixhe/laixhe-app/laixhe-server/protoapi"
	"github.com/laixhe/laixhe-app/laixhe-server/servers"
	"strconv"
	"time"
)

//====================================================

func WSLogin(w *servers.Request) {

	login := new(protoapi.LoginRequest)
	err := w.Message(login)
	if err != nil {
		fmt.Println("T------- GetMsg:", err)
		return
	}

	fmt.Printf("T---. login: %#v\n", login)

	err = w.Send(&protoapi.LoginResponse{Cmd:protoapi.CMD_C_LOGIN_REQUEST, Message: "--user_id--" + strconv.Itoa(int(time.Now().Unix()))})
	if err != nil {
		fmt.Println("T------- Send:", err)
	}

}
