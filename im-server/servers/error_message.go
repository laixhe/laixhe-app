package servers

import (
	"im-server/protoim"
)

// ErrorMessage 错误消息
func ErrorMessage(e protoim.E, msg ...string) *protoim.ErrorInfo {
	s := ""
	if len(msg) > 0 {
		s = msg[0]
	} else {
		s = e.String()
	}
	return &protoim.ErrorInfo{
		Code: e,
		Msg:  s,
	}
}

// EnCodeErrorMessage 编码错误消息
func EnCodeErrorMessage(e *protoim.ErrorInfo) ([]byte, error) {
	data, err := EnCode(protoim.CMD_C_ERROR, e)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// EnCodeError 编码错误消息
func EnCodeError(e protoim.E, msg ...string) ([]byte, error) {
	return EnCodeErrorMessage(ErrorMessage(e, msg...))
}
