package initialize

import (
	"fmt"

	"github.com/vuvanthanhtb/go-ecommerce-backend/global"
	"go.uber.org/zap"
)

func Run() {
	LoadConfig()
	fmt.Printf("Loading configuration mysql: %s\n", global.Config.MySQL.Username)
	InitLogger()
	global.Logger.Info("Config log OK.", zap.String("OK", "SUCCESS"))
	InitMysql()
	InitRedis()

	r := InitRouter()
	r.Run(":8082")
}
