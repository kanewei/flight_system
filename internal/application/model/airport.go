package model

type (
	CreateAirportRequest struct {
		Code      string `validate:"required" err_info:"code is required" json:"code"`
		Terminal  string `validate:"required" err_info:"terminal is required" json:"terminal"`
		City      string `validate:"required" err_info:"city is required" json:"city"`
		Name      string `validate:"required" err_info:"name is required" json:"name"`
		Available bool   `json:"available"`
	}
)

type (
	GetAirportRequest struct {
		Code string `json:"code"`
		City string `json:"city"`
	}
	GetAirportResponse struct {
		Code      string `json:"code"`
		Terminal  string `json:"terminal"`
		City      string `json:"city"`
		Name      string `json:"name"`
		Available bool   `json:"available"`
	}
)
