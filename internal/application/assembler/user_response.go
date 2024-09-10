package assembler

import "flight_system/internal/application/model"

type UserResponse struct{}

func NewUserResponse() *UserResponse {
	return new(UserResponse)
}

func (a *UserResponse) SignUpToResponse(id int64) *model.SignUpResponse {
	return &model.SignUpResponse{
		Id: id,
	}
}

func (a *UserResponse) LoginToResponse(id int64, name, token string) *model.LoginResponse {
	return &model.LoginResponse{
		Id:    id,
		Token: token,
		Name:  name,
	}
}
