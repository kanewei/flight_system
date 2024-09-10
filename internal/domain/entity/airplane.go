package entity

type Airplane struct {
	ID        int64   `json:"id"`
	Model     string  `json:"model"`
	Seats     []*Seat `json:"seats"`
	Available bool    `json:"available"`
}

type Seat struct {
	ID         string `json:"id"`
	Class      string `json:"class"`
	IsOverSold bool   `json:"is_over_sold"`
	Available  bool   `json:"available"`
}

func SeatToFlightSeat(seat *Seat) *FlightSeat {
	return &FlightSeat{
		SeatID:     seat.ID,
		Class:      seat.Class,
		IsOverSold: seat.IsOverSold,
		Available:  seat.Available,
	}
}

func SeatsToFlightSeats(seats []*Seat) []*FlightSeat {
	var response []*FlightSeat
	for _, seat := range seats {
		response = append(response, SeatToFlightSeat(seat))
	}
	return response
}
