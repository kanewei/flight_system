package entity

import (
	"strings"
	"time"
)

type Flight struct {
	ID            string       `json:"id"`
	Departure     string       `json:"departure"`
	Arrival       string       `json:"arrival"`
	DepartureTime time.Time    `json:"departure_time"`
	ArrivalTime   time.Time    `json:"arrival_time"`
	Price         string       `json:"price"`
	AirportCode   string       `json:"airport_code"`
	Airport       Airport      `json:"airport"`
	AirplaneID    int64        `json:"airplane_id"`
	Airplane      Airplane     `json:"airplane"`
	Seats         []FlightSeat `json:"seats"`
}

type FlightSeat struct {
	FlightID   string `json:"flight_id"`
	SeatID     string `json:"seat_id"`
	Class      string `json:"class"`
	IsOverSold bool   `json:"is_over_sold"`
	Available  bool   `json:"available"`
}

func SeatsToFlightTickets(flightId string, seats []*Seat) map[string]interface{} {
	ticketMap := make(map[string]interface{})

	var builder strings.Builder
	for _, seat := range seats {
		builder.WriteString(flightId)
		builder.WriteString(":")
		builder.WriteString(seat.ID)

		result := builder.String()
		ticketMap[result] = true
	}
	return ticketMap
}
