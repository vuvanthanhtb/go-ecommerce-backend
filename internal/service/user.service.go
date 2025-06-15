package service

import (
	"github.com/vuvanthanhtb/go-ecommerce-backend/internal/repository"
	"github.com/vuvanthanhtb/go-ecommerce-backend/packages/response"
)

type IUserService interface {
	Register(email string, purpose string) int
}

type userService struct {
	useRepository repository.IUserRepository
}

// Register implements IUserService.
func (us *userService) Register(email string, purpose string) int {
	if us.useRepository.GetUserByEmail(email) {
		return response.ErrCodeUserHasExists
	}
	return response.ErrCodeSuccess
}

func NewUserService(useRepository repository.IUserRepository) IUserService {
	return &userService{
		useRepository: useRepository,
	}
}
