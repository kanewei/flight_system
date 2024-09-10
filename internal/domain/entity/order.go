package entity

import "time"

type Order struct {
	ID        int64     `json:"id"`
	TicketID  string    `json:"ticket_id"`
	UserID    int64     `json:"user_id"`
	Price     string    `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}
