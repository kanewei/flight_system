package repository

import (
	"context"
	"flight_system/internal/domain/entity"
	"time"
)

type FlightRepository interface {
	CreateFlight(flight *entity.Flight) (*entity.Flight, error)
	GetFlightById(id string) (*entity.Flight, error)
	GetFlightByIdWithAvaliableSeats(id string) (*entity.Flight, error)
	GetFlightByIdAndSeatId(flightId string, seatId string) (*entity.Flight, error)
	SearchFlight(flight *entity.Flight, page int) ([]*entity.Flight, error)
}

type FlightSeatRepository interface {
	CreateFlightSeats(flightId string, flightSeats []*entity.FlightSeat) error
}

type FlightTicketCacheRepository interface {
	CreateFlightTickets(ctx context.Context, flightSeats map[string]interface{}, expireTime time.Duration) error
}
