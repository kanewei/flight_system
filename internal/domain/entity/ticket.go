package entity

type TicketStatus int

const (
	TicketStatusAvailable TicketStatus = iota
	TicketStatusOrdered
	TicketStatusBooked
)

type TicketOrder struct {
	ID       string `json:"id"`
	FlightID string `json:"flight_id"`
	SeatID   string `json:"seat_id"`
	UserID   int64  `json:"user_id"`
	Price    string `json:"price"`
}

type Ticket struct {
	ID       string       `json:"id"`
	FlightID string       `json:"flight_id"`
	SeatID   string       `json:"seat_id"`
	Flight   Flight       `json:"flight"`
	UserID   int64        `json:"user_id"`
	Price    string       `json:"price"`
	Status   TicketStatus `json:"status"`
}
