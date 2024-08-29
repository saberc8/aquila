package main

import (
	"aquila/global"
	"aquila/initialize"
	"go.uber.org/zap"
)

func main() {
	// 初始化配置
	global.AQUILA_VIPER = initialize.InitViper()
	// 初始化zap日志库
	global.AQUILA_LOG = initialize.InitZap()
	zap.ReplaceGlobals(global.AQUILA_LOG)

	global.AQUILA_LOG.Info("server run success on ", zap.String("zap_log", "zap_log"))
	initialize.InitServer()
}
