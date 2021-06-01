package servers

import (
	"im-server/protoim"
)

// ErrorMessage 错误消息
func ErrorMessage(e protoim.E, msg ...string) *protoim.ErrorInfo {
	msgStr := ""
	if len(msg) > 0 {
		msgStr = msg[0]
	}
	return &protoim.ErrorInfo{
		Code: e,
		Msg:  msgStr,
	}
}

// EnErrorMessage 编码错误消息
func EnErrorMessage(e *protoim.ErrorInfo) ([]byte, error) {
	data, err := EnCode(protoim.CMD_CError, e)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// ErrorMessageEnCode 错误消息-编码
func ErrorMessageEnCode(msg ...string) ([]byte, error) {
	em := ErrorMessage(protoim.E_EnCodeError, msg...)
	data, err := EnCode(protoim.CMD_CError, em)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// ErrorMessageDeCode 错误消息-解码
func ErrorMessageDeCode(msg ...string) ([]byte, error) {
	em := ErrorMessage(protoim.E_DeCodeError, msg...)
	data, err := EnCode(protoim.CMD_CError, em)
	if err != nil {
		return nil, err
	}
	return data, nil
}
