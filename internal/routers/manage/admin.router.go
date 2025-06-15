package manage

import "github.com/gin-gonic/gin"

type AdminnRouter struct {
}

func (pr *AdminnRouter) InitAdminRouter(routerGroup *gin.RouterGroup) {
	// public router
	adminRouterPublic := routerGroup.Group("/admin")
	{
		adminRouterPublic.POST("/login")
	}
	// private router
	userRouterPrivate := routerGroup.Group("/admin/user")
	// userRouterPrivate.Use(limiter())
	// userRouterPrivate.Use(Autth())
	// userRouterPrivate.Use(Permission())
	{
		userRouterPrivate.POST("/activate")
	}
}
