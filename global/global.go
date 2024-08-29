package global

import (
	"aquila/config"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var (
	AQUILA_CONFIG config.Configuration
	AQUILA_VIPER  *viper.Viper
	AQUILA_LOG    *zap.Logger
)
