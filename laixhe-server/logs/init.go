package logs

import (
	"github.com/laixhe/goutil/zap_log"
	"github.com/laixhe/laixhe-app/laixhe-server/config"
)

// 初始日志
func Init() {

	zap_log.Init(
		config.GetLog().Path,
		config.GetLog().Level,
		config.GetLog().MaxSize,
		config.GetLog().MaxBackup,
		config.GetLog().MaxAge,
		1)

}