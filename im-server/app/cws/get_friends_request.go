package cws

import (
	"fmt"

	"im-server/protoim"
	"im-server/servers"
)

// 获取好友列表-请求
func GetFriendsRequest(c *servers.Context) {
	req := new(protoim.GetFriendsRequest)
	err := c.ProtoBind(req)
	if err != nil {
		c.SendError(protoim.Error_E_PARAMETER, err.Error())
		return
	}

	users := make([]*protoim.UserInfo, 0)
	for _, v := range us {
		users = append(users, &protoim.UserInfo{
			UserId:   v.UserId,
			NickName: v.NickName,
		})
	}

	rsp := &protoim.GetFriendsResponse{
		Users: users,
	}

	e := c.Send(protoim.CMD_C_GET_FRIENDS_RESPONSE, rsp)
	fmt.Println("GetUserInfo - rsp:", rsp, "err:", e)
}
