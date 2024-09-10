package repository

import (
	"flight_system/internal/domain/entity"
)

type UserRepository interface {
	SignUp(user *entity.SignUp) (int64, error)
	Login(user *entity.Login) (*entity.User, error)
}
