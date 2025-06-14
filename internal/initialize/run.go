package initialize

import (
	"fmt"

	"github.com/vuvanthanhtb/go-ecommerce-backend/global"
)

func Run() {
	LoadConfig()
	fmt.Printf("Loading configuration mysql: %s\n", global.Config.MySQL.Username)
	InitLogger()
	InitMysql()
	InitRedis()

	r := InitRouter()
	r.Run(":8082")
}
