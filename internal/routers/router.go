package routers

// import (
// 	"github.com/gin-gonic/gin"
// 	"github.com/vuvanthanhtb/go-ecommerce-backend/internal/controller"
// 	"github.com/vuvanthanhtb/go-ecommerce-backend/internal/middlewares"
// )

// func NewRouter() *gin.Engine {
// 	r := gin.Default()

// 	// use the middlewwares
// 	r.Use(middlewares.AuthMiddleware())

// 	v1 := r.Group("/api")
// 	{
// 		v1.GET("/ping", controller.NewPongController().Pong)
// 		v1.GET("/info", controller.NewUserController().GetInfoUser)
// 	}
// 	return r
// }
