package model

import "time"

type (
	CreateFlightRequest struct {
		ID            string    `json:"id"`
		Departure     string    `json:"departure"`
		Arrival       string    `json:"arrival"`
		DepartureTime time.Time `json:"departure_time"`
		ArrivalTime   time.Time `json:"arrival_time"`
		Price         string    `json:"price"`
		AirportCode   string    `json:"airport_code"`
		AirplaneID    int64     `json:"airplane_id"`
	}
	CreateFlightResponse struct {
		ID string `json:"id"`
	}
)

type (
	GetFlightRequest struct {
		ID string `json:"id"`
	}
	GetFlightResponse struct {
		ID                     string                   `json:"id"`
		Departure              string                   `json:"departure"`
		Arrival                string                   `json:"arrival"`
		DepartureTime          time.Time                `json:"departure_time"`
		ArrivalTime            time.Time                `json:"arrival_time"`
		Price                  string                   `json:"price"`
		AirportCode            string                   `json:"airport_code"`
		Airport                GetAirportResponse       `json:"airport"`
		AirplaneID             int64                    `json:"airplane_id"`
		Airplane               GetAirplaneResponse      `json:"airplane"`
		Seats                  []GetFlightSeatsResponse `json:"seats"`
		RemainingOverSoldSeats bool                     `json:"remaining_over_sold_seats"`
	}
)

type (
	SearchFlightRequest struct {
		Departure     string    `json:"departure"`
		Arrival       string    `json:"arrival"`
		DepartureTime time.Time `json:"departure_time"`
		ArrivalTime   time.Time `json:"arrival_time"`
		Page          int       `validate:"gt=0" json:"page"`
	}
	SearchFlightResponse struct {
		ID            string                   `json:"id"`
		Departure     string                   `json:"departure"`
		Arrival       string                   `json:"arrival"`
		DepartureTime time.Time                `json:"departure_time"`
		ArrivalTime   time.Time                `json:"arrival_time"`
		Price         string                   `json:"price"`
		AirportCode   string                   `json:"airport_code"`
		Airport       GetAirportResponse       `json:"airport"`
		AirplaneID    int64                    `json:"airplane_id"`
		Airplane      GetAirplaneResponse      `json:"airplane"`
		Seats         []GetFlightSeatsResponse `json:"seats"`
	}
)

type (
	GetFlightSeatsResponse struct {
		FlightID   string `json:"flight_id"`
		ID         string `json:"id"`
		Class      string `json:"class"`
		IsOverSold bool   `json:"is_over_sold"`
		Available  bool   `json:"available"`
	}
)
