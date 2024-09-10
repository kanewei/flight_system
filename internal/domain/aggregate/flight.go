package aggregate

import (
	"context"
	"flight_system/internal/domain/entity"
	"flight_system/internal/domain/repository"
	flightTicketRepoImpl "flight_system/internal/infrastructure/repository/impl/flight/cache"
	flightRepoImpl "flight_system/internal/infrastructure/repository/impl/flight/database"
	"time"
)

type FlightAggrate interface {
	CreateFlight(ctx context.Context, flight *entity.Flight) (string, error)
	GetFlightById(id string) (*entity.Flight, error)
	SearchFlight(flight *entity.Flight, page int) ([]*entity.Flight, error)
}

type flightAggrate struct {
	flightRepo            repository.FlightRepository
	flightSeatRepo        repository.FlightSeatRepository
	flightTicketCacheRepo repository.FlightTicketCacheRepository
}

func NewFlightAggrate() *flightAggrate {
	return &flightAggrate{
		flightRepo:            flightRepoImpl.NewFlightRepoImpl(),
		flightSeatRepo:        flightRepoImpl.NewFlightSeatRepoImpl(),
		flightTicketCacheRepo: flightTicketRepoImpl.NewFlightTicketCacheRepoImpl(),
	}
}

func (f *flightAggrate) CreateFlight(ctx context.Context, flight *entity.Flight) (string, error) {
	// Todo check if airplane is available
	// Todo check if airport is available

	createdFlight, err := f.flightRepo.CreateFlight(flight)
	if err != nil {
		return "", err
	}

	if err := f.flightSeatRepo.CreateFlightSeats(createdFlight.ID, entity.SeatsToFlightSeats(createdFlight.Airplane.Seats)); err != nil {
		return "", err
	}

	expireTime := time.Now().Sub(createdFlight.DepartureTime)
	if err := f.flightTicketCacheRepo.CreateFlightTickets(ctx, entity.SeatsToFlightTickets(createdFlight.ID, createdFlight.Airplane.Seats), expireTime); err != nil {
		return "", err
	}

	return createdFlight.ID, nil
}

func (f *flightAggrate) GetFlightById(id string) (*entity.Flight, error) {
	return f.flightRepo.GetFlightById(id)
}

func (f *flightAggrate) SearchFlight(flight *entity.Flight, page int) ([]*entity.Flight, error) {
	return f.flightRepo.SearchFlight(flight, page)
}
