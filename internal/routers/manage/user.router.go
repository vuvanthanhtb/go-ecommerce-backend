package manage

import "github.com/gin-gonic/gin"

type UserRouter struct {
}

func (pr *UserRouter) InitUserRouter(routerGroup *gin.RouterGroup) {
	// private router
	userRouterPrivate := routerGroup.Group("/admin/user")
	// userRouterPrivate.Use(limiter())
	// userRouterPrivate.Use(Autth())
	// userRouterPrivate.Use(Permission())
	{
		userRouterPrivate.POST("/activate-user")
	}
}
