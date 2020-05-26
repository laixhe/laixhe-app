package ws

import "encoding/json"

// 通用请求指令
type WSCmd struct {
	Cmd uint `json:"cmd"` // 请求指令
}

// 通用数据接口
type DataInterface interface {
	Encode() ([]byte, error) // 编码
}

// 通用数据格式接口
type WsDataInterface interface {
	GetCmd() uint // 指令
}

// 通用数据格式
type WsData struct {
	Cmd  uint        `json:"cmd"`            // 指令
	Data interface{} `json:"data,omitempty"` // 数据
}

func (this *WsData) Encode() ([]byte, error) {

	js, err := json.Marshal(this)
	if err != nil {
		return nil, err
	}

	return js, nil
}

func NewWsData(data WsDataInterface) *WsData {
	return &WsData{Cmd: data.GetCmd(), Data: data}
}
