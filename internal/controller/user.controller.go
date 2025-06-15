package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/vuvanthanhtb/go-ecommerce-backend/internal/service"
	"github.com/vuvanthanhtb/go-ecommerce-backend/packages/response"
)

type UserController struct {
	userService service.IUserService
}

// Register implements IUserController.
func (uc *UserController) Register(ctx *gin.Context) {
	result := uc.userService.Register("", "")
	response.SuccessResponse(ctx, result, nil)
}

func NewUserController(userService service.IUserService) *UserController {
	return &UserController{
		userService: userService,
	}
}
