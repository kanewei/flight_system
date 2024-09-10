package assembler

import (
	"flight_system/internal/application/model"
	"flight_system/internal/domain/entity"
)

type UserResponse struct{}

func NewUserResponse() *UserResponse {
	return new(UserResponse)
}

func (a *UserResponse) SignUpToResponse(id int64) *model.SignUpResponse {
	return &model.SignUpResponse{
		Id: id,
	}
}

func (a *UserResponse) LoginToResponse(user *entity.User, token string) *model.LoginResponse {
	return &model.LoginResponse{
		Id:    user.ID,
		Token: token,
		Name:  user.Name,
	}
}
