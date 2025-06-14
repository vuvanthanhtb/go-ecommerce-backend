package global

import (
	"github.com/vuvanthanhtb/go-ecommerce-backend/packages/logger"
	"github.com/vuvanthanhtb/go-ecommerce-backend/packages/setting"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Mdb    *gorm.DB
)
