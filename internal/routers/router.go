package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/vuvanthanhtb/go-ecommerce-backend/internal/controller"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api")
	{
		v1.GET("/ping", controller.NewPongController().Pong)
		v1.GET("/info", controller.NewUserController().GetInfoUser)
	}
	return r
}
