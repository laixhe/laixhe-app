package cache

import (
	"fmt"

	"im-server/database/redisx"
)

const (
	ptsKey = "pts_%s" // 当前会话的序列ID (自增的)
)

// GetPtsID 获取每个会话的 pts
func GetPtsID(dialogId string) (int64, error) {
	key := fmt.Sprintf(ptsKey, dialogId)
	return redisx.Incr(key)
}
