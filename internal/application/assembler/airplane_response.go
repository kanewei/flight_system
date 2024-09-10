package assembler

import (
	"flight_system/internal/application/model"
	"flight_system/internal/domain/entity"
)

type AirplaneResponse struct{}

func NewAirplaneResponse() *AirplaneResponse {
	return new(AirplaneResponse)
}

func (a *AirplaneResponse) CreateToResponse(id int64) *model.CreateAirplaneResponse {
	return &model.CreateAirplaneResponse{
		ID: id,
	}
}

func airplaneToResponse(airplane *entity.Airplane) *model.GetAirplaneResponse {
	return &model.GetAirplaneResponse{
		ID:        airplane.ID,
		Model:     airplane.Model,
		Seats:     seatsToResponse(airplane.Seats),
		Available: airplane.Available,
	}
}

func (a *AirplaneResponse) GetByIdToResponse(airplane *entity.Airplane) *model.GetAirplaneResponse {
	return airplaneToResponse(airplane)
}

func (a *AirplaneResponse) GetByModelToResponse(airplanes []*entity.Airplane) []*model.GetAirplaneResponse {
	var response []*model.GetAirplaneResponse
	for _, airplane := range airplanes {
		response = append(response, airplaneToResponse(airplane))
	}
	return response
}

func seatToResponse(seat *entity.Seat) *model.GetSeatsResponse {
	return &model.GetSeatsResponse{
		ID:    seat.ID,
		Class: seat.Class,
	}
}

func seatsToResponse(seats []*entity.Seat) []*model.GetSeatsResponse {
	var response []*model.GetSeatsResponse
	for _, seat := range seats {
		response = append(response, seatToResponse(seat))
	}
	return response
}
