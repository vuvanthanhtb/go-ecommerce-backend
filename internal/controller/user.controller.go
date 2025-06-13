package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/vuvanthanhtb/go-ecommerce-backend/internal/service"
	"github.com/vuvanthanhtb/go-ecommerce-backend/packages/response"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{userService: service.NewUserService()}
}

func (uc *UserController) GetInfoUser(ctx *gin.Context) {
	// response.SuccessResponse(ctx, 2001, "success", uc.userService.GetInfoUser())
	response.ErrorResponse(ctx, 2003)
}
