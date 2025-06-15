package user

import "github.com/gin-gonic/gin"

type UserRouter struct {
}

func (pr *UserRouter) InitUserRouter(routerGroup *gin.RouterGroup) {
	// public router
	userRouterPublic := routerGroup.Group("/user")
	{
		userRouterPublic.POST("/register")
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
