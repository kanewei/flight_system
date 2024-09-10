package aggregate

import (
	"context"
	"flight_system/internal/domain/entity"
	"flight_system/internal/domain/repository"
	"flight_system/internal/global"
	flightCacheRepoImpl "flight_system/internal/infrastructure/repository/impl/flight/cache"
	flightTicketRepoImpl "flight_system/internal/infrastructure/repository/impl/flight/cache"
	flightRepoImpl "flight_system/internal/infrastructure/repository/impl/flight/database"
	"time"
)

type FlightAggrate interface {
	CreateFlight(ctx context.Context, flight *entity.Flight) (string, error)
	GetFlightById(id string) (*entity.Flight, error)
	SearchFlight(ctx context.Context, flight *entity.Flight, page int) ([]*entity.Flight, error)
}

type flightAggrate struct {
	flightRepo            repository.FlightRepository
	flightSeatRepo        repository.FlightSeatRepository
	flightCacheRepo       repository.FlightCacheRepository
	flightTicketCacheRepo repository.FlightTicketCacheRepository
}

func NewFlightAggrate() *flightAggrate {
	return &flightAggrate{
		flightRepo:            flightRepoImpl.NewFlightRepoImpl(),
		flightSeatRepo:        flightRepoImpl.NewFlightSeatRepoImpl(),
		flightCacheRepo:       flightCacheRepoImpl.NewFlightCacheRepoImpl(),
		flightTicketCacheRepo: flightTicketRepoImpl.NewFlightTicketCacheRepoImpl(),
	}
}

func (f *flightAggrate) CreateFlight(ctx context.Context, flight *entity.Flight) (string, error) {
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

func (f *flightAggrate) SearchFlight(ctx context.Context, flight *entity.Flight, page int) ([]*entity.Flight, error) {
	flights, err := f.flightCacheRepo.GetSearchFlight(ctx, flight, page)
	if err != nil {
		global.Log.Errorf("Failed to get search flight from cache, error: %v", err)
	}
	if len(flights) > 0 {
		return flights, nil
	}

	flights, err = f.flightRepo.SearchFlight(flight, page)
	if err != nil {
		return nil, err
	}

	err = f.flightCacheRepo.SetSearchFlight(ctx, flight, flights)
	if err != nil {
		global.Log.Errorf("Failed to set search flight to cache, error: %v", err)
	}

	return flights, nil
}
