package model

type (
	CreateStaffRequest struct {
		Email    string `validate:"required,email" json:"email"`
		Name     string `validate:"required" json:"Name"`
		Password string `validate:"required" json:"password"`
	}
	CreateStaffResponse struct {
		ID int64 `json:"id"`
	}
)

type (
	StaffLoginRequest struct {
		Email    string `validate:"required,email" json:"email"`
		Password string `validate:"required" json:"password"`
	}
	StaffLoginResponse struct {
		ID    int64  `json:"id"`
		Token string `json:"token"`
		Name  string `json:"name"`
	}
)
