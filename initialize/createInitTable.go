package initialize

import (
	"aquila/api/system"
	"aquila/global"
	"go.uber.org/zap"
	"os"
)

func CreateInitTable() {
	db := global.AquilaDb
	err := db.AutoMigrate(
		system.User{},
	)
	if err != nil {
		global.AquilaLog.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.AquilaLog.Info("register table success")
}
