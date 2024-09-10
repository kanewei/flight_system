package model

type (
	SignUpRequest struct {
		Email    string `validate:"required" err_info:"email is required" json:"email"`
		Password string `validate:"required" err_info:"password is required" json:"password"`
		Name     string `validate:"required" err_info:"name is required" json:"name"`
	}
	SignUpResponse struct {
		Id int64 `json:"id"`
	}
)

type (
	LoginRequest struct {
		Email    string `validate:"required" err_info:"email is required" json:"email"`
		Password string `validate:"required" err_info:"password is required" json:"password"`
	}
	LoginResponse struct {
		Id    int64  `json:"id"`
		Token string `json:"token"`
		Name  string `json:"name"`
	}
)
