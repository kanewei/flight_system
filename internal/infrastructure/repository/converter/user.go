package converter

import (
	"flight_system/internal/domain/entity"
	"flight_system/internal/infrastructure/repository/po"
)

func SignUpToPo(user *entity.SignUp) *po.User {
	return &po.User{
		Password: user.Password,
		Email:    user.Email,
		Name:     user.Name,
	}
}

func LoginToPo(user *entity.Login) *po.User {
	return &po.User{
		Password: user.Password,
		Email:    user.Email,
	}
}

func UserPoToEntity(po *po.User) *entity.User {
	return &entity.User{
		ID:    po.ID,
		Email: po.Email,
		Name:  po.Name,
	}
}
