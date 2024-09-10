package repository

import "flight_system/internal/domain/entity"

type AirportRepository interface {
	CreateAirport(airport *entity.Airport) error
	GetAirportByCode(code string) (*entity.Airport, error)
	GetAirportByCity(city string) ([]*entity.Airport, error)
}
