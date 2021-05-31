package utils

import (
	"fmt"

	"im-server/protoim"
)

// GetUserKey 用户平台连接key
func GetUserKey(appOsId protoim.AppOs, userId string) string {
	return fmt.Sprintf("%d_%s", appOsId, userId)
}
