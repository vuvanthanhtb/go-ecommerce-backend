package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/vuvanthanhtb/go-ecommerce-backend/global"
	"github.com/vuvanthanhtb/go-ecommerce-backend/internal/routers"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	if global.Config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}

	manageRouter := routers.RouterGroupApplication.Manage
	userRouter := routers.RouterGroupApplication.User

	MainGroup := r.Group("/v1")
	{
		MainGroup.GET("/check-status")
	}
	{
		manageRouter.InitAdminRouter(MainGroup)
		manageRouter.InitUserRouter(MainGroup)
	}
	{
		userRouter.InitProductRouter(MainGroup)
		userRouter.InitUserRouter(MainGroup)
	}
	return r
}
