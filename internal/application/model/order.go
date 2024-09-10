package model

type (
	CreateOrderRequest struct {
		TicketID string `validate:"required" err_info:"ticket_id is required" json:"ticket_id"`
		UserID   string `validate:"required" err_info:"user_id is required" json:"user_id"`
	}
	CreateOrderResponse struct {
		ID int64 `json:"id"`
	}
)
