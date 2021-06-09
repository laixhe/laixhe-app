package utils

import (
	"fmt"

	"im-server/protoim"
)

// GetUserKey 用户平台连接key
func GetUserKey(osID protoim.AppOs, userID string) string {
	return fmt.Sprintf("%d_%s", osID, userID)
}

// IsAppOs 是否是已存在的平台
func IsAppOs(osID protoim.AppOs) bool {
	if protoim.AppOs_WEB == osID {
		return true
	}
	if protoim.AppOs_IOS == osID {
		return true
	}
	if protoim.AppOs_ANDROID == osID {
		return true
	}
	if protoim.AppOs_PC == osID {
		return true
	}
	return false
}
