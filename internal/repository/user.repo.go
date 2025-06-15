package repository

// type UserRepo struct{}

// func NewUserRepo() *UserRepo {
// 	return &UserRepo{}
// }

// func (ur *UserRepo) GetInfoUser() string {
// 	return "thanhvv"
// }

type IUserRepository interface {
	GetUserByEmail(email string) bool
}

type userReposiory struct {
}

// GetUserByEmail implements IUserRepository.
func (u *userReposiory) GetUserByEmail(email string) bool {
	return true
}

func NewUserRepository() IUserRepository {
	return &userReposiory{}
}
