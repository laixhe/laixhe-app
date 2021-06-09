package servers

import (
	"google.golang.org/protobuf/proto"

	"im-server/protoim"
)

// EnCode 编码
func EnCode(cmd protoim.CMD, data proto.Message) ([]byte, error) {
	protoData, err := proto.Marshal(data)
	if err != nil {
		return nil, err
	}
	messageBase := &protoim.MessageBase{
		Cmd:  cmd,
		Data: protoData,
	}
	protoBase, err := proto.Marshal(messageBase)
	if err != nil {
		return nil, err
	}
	return protoBase, nil
}

// EnCode 编码 CMD
func EnCodeCmd(cmd protoim.CMD) ([]byte, error) {
	messageBase := &protoim.MessageBase{
		Cmd: cmd,
	}
	protoBase, err := proto.Marshal(messageBase)
	if err != nil {
		return nil, err
	}
	return protoBase, nil
}

// DeCodeMessageBase 解码
func DeCodeMessageBase(data []byte) (*protoim.MessageBase, *protoim.ErrorInfo) {
	messageBase := new(protoim.MessageBase)
	err := proto.Unmarshal(data, messageBase)
	if err != nil {
		return nil, ErrorMessage(protoim.E_DECODE_ERROR, err.Error())
	}
	return messageBase, nil
}
