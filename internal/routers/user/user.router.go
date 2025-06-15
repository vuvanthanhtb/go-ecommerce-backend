package user

import (
	"github.com/gin-gonic/gin"
	"github.com/vuvanthanhtb/go-ecommerce-backend/internal/wire"
)

type UserRouter struct {
}

func (pr *UserRouter) InitUserRouter(routerGroup *gin.RouterGroup) {
	userController, _ := wire.InitUserRouterHandler()
	// public router
	userRouterPublic := routerGroup.Group("/user")
	{
		userRouterPublic.POST("/register", userController.Register)
		userRouterPublic.POST("/send-otp")
	}
	// private router
	userRouterPrivate := routerGroup.Group("/user")
	// userRouterPrivate.Use(limiter())
	// userRouterPrivate.Use(Autth())
	// userRouterPrivate.Use(Permission())
	{
		userRouterPrivate.GET("/get-info")
	}
}
