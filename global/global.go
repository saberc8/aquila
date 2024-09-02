package global

import (
	"aquila/config"
	"github.com/go-redis/redis/v8"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	AquilaConfig config.Configuration
	AquilaViper  *viper.Viper
	AquilaLog    *zap.Logger
	AquilaDb     *gorm.DB // DB 数据库连接
	AquilaRedis  *redis.Client
)
