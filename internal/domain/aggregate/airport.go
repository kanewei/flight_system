package aggregate

import (
	"flight_system/internal/domain/entity"
	"flight_system/internal/domain/repository"
	airportRepoImpl "flight_system/internal/infrastructure/repository/impl/airport/database"
)

type AirportAggrate interface {
	CreateAirport(airport *entity.Airport) error
	GetAirportByCode(code string) (*entity.Airport, error)
	GetAirportByCity(city string) ([]*entity.Airport, error)
}

type airportAggrate struct {
	airportRepo repository.AirportRepository
}

func NewAirportAggrate() *airportAggrate {
	return &airportAggrate{
		airportRepo: airportRepoImpl.NewAirportRepoImpl(),
	}
}

func (a *airportAggrate) CreateAirport(airport *entity.Airport) error {
	return a.airportRepo.CreateAirport(airport)
}

func (a *airportAggrate) GetAirportByCode(code string) (*entity.Airport, error) {
	return a.airportRepo.GetAirportByCode(code)
}

func (a *airportAggrate) GetAirportByCity(city string) ([]*entity.Airport, error) {
	return a.airportRepo.GetAirportByCity(city)
}
