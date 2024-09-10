package assembler

import (
	"flight_system/internal/application/model"
	"flight_system/internal/domain/entity"

	"github.com/jinzhu/copier"
)

type UserRequest struct{}

func NewUserRequest() *UserRequest {
	return new(UserRequest)
}

func (u *UserRequest) UserSignUpRequestToEntity(req *model.SignUpRequest) *entity.SignUp {
	var entity entity.SignUp
	if err := copier.Copy(&entity, &req); err != nil {
		panic(any(err))
	}

	return &entity
}

func (u *UserRequest) UserLoginRequestToEntity(req *model.LoginRequest) *entity.Login {
	var entity entity.Login
	if err := copier.Copy(&entity, &req); err != nil {
		panic(any(err))
	}

	return &entity
}
