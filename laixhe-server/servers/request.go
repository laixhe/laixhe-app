package servers

import (
	"github.com/laixhe/laixhe-app/laixhe-server/protoapi"
	"google.golang.org/protobuf/proto"
)

// 数据包的请求
type Request struct {
	conn *Client      // 客户端链接
	msg  []byte       // 客户端请求的数据
	cmd  protoapi.CMD // 客户端请求指令
}

func NewRequest(conn *Client, msg []byte, cmd protoapi.CMD) *Request {
	return &Request{
		conn: conn,
		msg:  msg,
		cmd:  cmd,
	}
}

// 客户端请求指令
func (this *Request) Cmd() protoapi.CMD {
	return this.cmd
}

// 客户端请求的数据
func (this *Request) Message(data proto.Message) error {
	return proto.Unmarshal(this.msg, data)
}

// 客户端发送消息
func (this *Request) Send(data proto.Message) error {

	code, err := proto.Marshal(data)
	if err != nil {
		return err
	}

	return this.conn.SendClient(code)
}

// 保存在线用户
func (this *Request) SaveUser(appId AppOS, userId string) {
	UserManager.SaveUser(NewClientUser(this.conn, appId, userId))
}
