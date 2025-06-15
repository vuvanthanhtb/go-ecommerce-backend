package global

import (
	"github.com/redis/go-redis/v9"
	"github.com/vuvanthanhtb/go-ecommerce-backend/packages/logger"
	"github.com/vuvanthanhtb/go-ecommerce-backend/packages/setting"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Mdb    *gorm.DB
	Rdb    *redis.Client
)
