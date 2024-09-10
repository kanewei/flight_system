package model

type (
	CreateAirplaneRequest struct {
		Model string                `validate:"required" err_info:"model is required" json:"model"` // airplane model (e.g. Boeing 737, Airbus A320)
		Seats []*CreateSeatsRequest `json:"seats"`
	}
	CreateAirplaneResponse struct {
		ID int64 `json:"id"`
	}
	CreateSeatsRequest struct {
		ID         string `json:"id"`    // seat code (e.g. 1A, 2B, 3C)
		Class      string `json:"class"` // seat class (e.g. economy, business, first)
		IsOverSold bool   `json:"is_over_sold"`
		Available  bool   `json:"available"`
	}
)

type (
	GetAirplaneRequest struct {
		ID    int64  `json:"id"`
		Model string `json:"model"`
	}
	GetAirplaneResponse struct {
		ID        int64               `json:"id"`
		Model     string              `json:"model"`
		Seats     []*GetSeatsResponse `json:"seats"`
		Available bool                `json:"available"`
	}
	GetSeatsResponse struct {
		ID        string `json:"id"`
		Class     string `json:"class"`
		Available bool   `json:"available"`
	}
)
