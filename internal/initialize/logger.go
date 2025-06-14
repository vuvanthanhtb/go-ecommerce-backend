package initialize

import (
	"github.com/vuvanthanhtb/go-ecommerce-backend/global"
	"github.com/vuvanthanhtb/go-ecommerce-backend/packages/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}
