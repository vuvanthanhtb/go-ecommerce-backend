package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vuvanthanhtb/go-ecommerce-backend/internal/service"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{userService: service.NewUserService()}
}

func (uc *UserController) GetInfoUser(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": uc.userService.GetInfoUser()})
}
