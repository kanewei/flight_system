package aggregate

import (
	"flight_system/internal/domain/entity"
	"flight_system/internal/domain/repository"
	userRepoImpl "flight_system/internal/infrastructure/repository/impl/user/database"
)

type UserAggrate interface {
	SignUp(user *entity.SignUp) (int64, error)
	Login(user *entity.Login) (*entity.User, error)
}

type userAggrate struct {
	userRepoImpl repository.UserRepository
}

func NewUserAggrate() *userAggrate {
	return &userAggrate{
		userRepoImpl: userRepoImpl.NewUserRepoImpl(),
	}
}

func (u *userAggrate) SignUp(user *entity.SignUp) (int64, error) {
	return u.userRepoImpl.SignUp(user)
}

func (u *userAggrate) Login(user *entity.Login) (*entity.User, error) {
	return u.userRepoImpl.Login(user)
}
