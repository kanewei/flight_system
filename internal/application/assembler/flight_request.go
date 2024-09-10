package assembler

import (
	"flight_system/internal/application/model"
	"flight_system/internal/domain/entity"

	"github.com/jinzhu/copier"
)

type FlightRequest struct{}

func NewFlightRequest() *FlightRequest {
	return new(FlightRequest)
}

func (f *FlightRequest) FlightRequestToEntity(req *model.CreateFlightRequest) *entity.Flight {
	var entity entity.Flight
	if err := copier.Copy(&entity, &req); err != nil {
		panic(any(err))
	}

	return &entity
}

func (f *FlightRequest) GetFlightRequestToEntity(req *model.GetFlightRequest) *entity.Flight {
	var entity entity.Flight
	if err := copier.Copy(&entity, &req); err != nil {
		panic(any(err))
	}

	return &entity
}

func (f *FlightRequest) SearchFlightRequestToEntity(req *model.SearchFlightRequest) *entity.Flight {
	var entity entity.Flight
	if err := copier.Copy(&entity, &req); err != nil {
		panic(any(err))
	}

	return &entity
}
