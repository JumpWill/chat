package global

import (
	"chat/schema"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
	"go.uber.org/zap"
)

var (
	// config
	Settings schema.ServerConfig
	// db con
	Db *gorm.DB
	// redis con
	Redis *redis.Client
	// log
	Log *zap.Logger
)
