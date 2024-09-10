package service

import (
	"flight_system/internal/application/assembler"
	"flight_system/internal/application/model"
	"flight_system/internal/domain/aggregate"
)

type UserService struct {
	userRequest  *assembler.UserRequest
	userResponse *assembler.UserResponse
	userAggrate  aggregate.UserAggrate
}

func NewUserService() *UserService {
	return &UserService{
		userRequest:  assembler.NewUserRequest(),
		userResponse: assembler.NewUserResponse(),
		userAggrate:  aggregate.NewUserAggrate(),
	}
}

func (u *UserService) SignUp(req *model.SignUpRequest) (*model.SignUpResponse, error) {
	user := u.userRequest.UserSignUpRequestToEntity(req)
	id, err := u.userAggrate.SignUp(user)
	if err != nil {
		return nil, err
	}
	return u.userResponse.SignUpToResponse(id), nil
}

func (u *UserService) Login(req *model.LoginRequest) (*model.LoginResponse, error) {
	login := u.userRequest.UserLoginRequestToEntity(req)
	user, err := u.userAggrate.Login(login)
	if err != nil {
		return nil, err
	}

	// Todo generate token
	token := ""

	return u.userResponse.LoginToResponse(user.ID, user.Name, token), nil
}
