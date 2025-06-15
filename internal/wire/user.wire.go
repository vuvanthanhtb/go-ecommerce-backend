//go:build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/vuvanthanhtb/go-ecommerce-backend/internal/controller"
	"github.com/vuvanthanhtb/go-ecommerce-backend/internal/repository"
	"github.com/vuvanthanhtb/go-ecommerce-backend/internal/service"
)

func InitUserRouterHandler() (*controller.UserController, error) {
	wire.Build(
		repository.NewUserRepository,
		service.NewUserService,
		controller.NewUserController,
	)

	return new(controller.UserController), nil
}
