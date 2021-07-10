package check

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"im-server/app/cache"
	"im-server/protoim"
	"im-server/utils/getid"
)

// ChatType 聊天类型是否存在
func ChatType(chatTypeID protoim.ChatType) bool {
	_, ok := protoim.ChatType_name[int32(chatTypeID)]
	return ok
}

// MessageType 消息类型是否存在
func MessageType(messageTypeID protoim.MessageType) bool {
	_, ok := protoim.MessageType_name[int32(messageTypeID)]
	return ok
}

// MessageSendRequest 检查发送消息(客户端发送-服务端接收)并生成相关ID
func MessageSendRequest(m *protoim.MessageSendRequest) error {
	if m.LocalMessageId == "" {
		return errors.New("no local_message_id")
	}
	if m.Content == nil {
		return errors.New("no content")
	}
	if m.Content.FromId == "" {
		return errors.New("no from_id")
	}
	if m.Content.ToId == "" {
		return errors.New("no to_id")
	}
	if !ChatType(m.Content.ChatTypeId) {
		return fmt.Errorf("there is no such chat_type_id:%d", m.Content.ChatTypeId)
	}
	if !MessageType(m.Content.MessageTypeId) {
		return fmt.Errorf("there is no such message_type_id:%d", m.Content.MessageTypeId)
	}
	dialogID := getid.DialogID(m.Content.FromId, m.Content.ToId, m.Content.ChatTypeId)
	if m.Content.DialogId == "" {
		m.Content.DialogId = dialogID
	} else {
		if m.Content.DialogId != dialogID {
			return errors.New("disaccord dialog_id")
		}
	}
	m.Content.MessageId = getid.MessageID()
	if m.Content.MessageTypeId == protoim.MessageType_M_TEXT {
		if m.Content.TextContent == "" {
			return errors.New("no text_content")
		}
		m.Content.TextContent = strings.TrimSpace(m.Content.TextContent)
		if m.Content.TextContent == "" {
			return errors.New("no text_content")
		}
	}
	if m.Content.MessageTypeId == protoim.MessageType_M_IMAGE ||
		m.Content.MessageTypeId == protoim.MessageType_M_FILE ||
		m.Content.MessageTypeId == protoim.MessageType_M_VOICE ||
		m.Content.MessageTypeId == protoim.MessageType_M_VIDEO {
		if m.Content.Media == nil {
			return errors.New("no media")
		}
	}
	pts, err := cache.GetPtsID(m.Content.DialogId)
	if err != nil {
		return errors.New("pts build failed:"+err.Error())
	}
	m.Content.Pts = pts
	m.Content.LatestTime = time.Now().Unix()
	return nil
}

