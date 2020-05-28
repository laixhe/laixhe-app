package servers

import "encoding/json"

// 数据包的请求
type Request struct {
	conn *Client // 客户端链接
	msg  []byte  // 客户端请求的数据
	cmd  uint    // 客户端请求指令
}

func NewRequest(conn *Client, msg []byte, cmd uint) *Request {
	return &Request{
		conn: conn,
		msg:  msg,
		cmd:  cmd,
	}
}

// 客户端请求指令
func (this *Request) Cmd() uint {
	return this.cmd
}

// 客户端请求的数据
func (this *Request) Message(data WsDataInterface) (*WsData, error) {

	wsData := &WsData{Data: data}

	err := json.Unmarshal(this.msg, wsData)
	if err != nil {
		return nil, err
	}

	return wsData, nil
}

// 客户端发送消息
func (this *Request) Send(data WsDataInterface) error {

	// 数据打包序列化
	wsData := NewWsData(data)
	code, err := wsData.Encode()
	if err != nil {
		return err
	}

	return this.conn.SendClient(code)
}

// 保存在线用户
func (this *Request) SaveUser(appId AppOS, userId string) {
	UserManager.SaveUser(NewClientUser(this.conn, appId, userId))
}
