package getid

import (
	"fmt"

	"github.com/laixhe/goutil/utils"

	"im-server/protoim"
)

// MessageID 获取消息id (以: m_ 开头)
func MessageID() string {
	return fmt.Sprintf("m_%d", utils.GetSnowFlakeID())
}

// DialogID 获取会话ID (以: d_ 开头) ( d_{user_id}_{chat_type}_{to_id} )
func DialogID(fromId string, toId string, chatTypeId protoim.ChatType) string {
	return fmt.Sprintf("d_%s_%d_%s", fromId, chatTypeId, toId)
}
