package model

type (
	CreateTicketOrderRequest struct {
		FlightID string `validate:"required" json:"flight_id"`
		SeatID   string `json:"seat_id"`
		UserID   int64  `validate:"required" json:"user_id"`
	}
	CreateTicketOrderResponse struct {
		ID    string `json:"id"`
		Price string `json:"price"`
	}
	CreateTicketRequest struct {
		FlightID string `json:"flight_id"`
		SeatID   string `json:"seat_id"`
		UserID   int64  `json:"user_id"`
	}
	CreateTicketResponse struct {
		ID string `json:"id"`
	}
)

type (
	GetTicketRequest struct {
		ID string `json:"id"`
	}
	GetTicketResponse struct {
		ID     string            `json:"id"`
		Flight GetFlightResponse `json:"flight"`
	}
)

type (
	GetUserTicketRequest struct {
		UserID int64 `json:"user_id"`
	}
	GetUserTicketResponse struct {
		ID     string            `json:"id"`
		Flight GetFlightResponse `json:"flight"`
	}
)
