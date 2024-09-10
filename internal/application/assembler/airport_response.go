package assembler

import (
	"flight_system/internal/application/model"
	"flight_system/internal/domain/entity"
)

type AirportResponse struct{}

func NewAirportResponse() *AirportResponse {
	return new(AirportResponse)
}

func airportToResponse(airport *entity.Airport) *model.GetAirportResponse {
	return &model.GetAirportResponse{
		Code:      airport.Code,
		Terminal:  airport.Terminal,
		City:      airport.City,
		Name:      airport.Name,
		Available: airport.Available,
	}
}

func (a *AirportResponse) GetByCodeToResponse(entity *entity.Airport) *model.GetAirportResponse {
	return airportToResponse(entity)
}

func (a *AirportResponse) GetByCityToResponse(entities []*entity.Airport) []*model.GetAirportResponse {
	var airports []*model.GetAirportResponse
	for _, entity := range entities {
		airports = append(airports, airportToResponse(entity))
	}
	return airports
}
