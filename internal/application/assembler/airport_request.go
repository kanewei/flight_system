package assembler

import (
	"flight_system/internal/application/model"
	"flight_system/internal/domain/entity"

	"github.com/jinzhu/copier"
)

type AirportRequest struct{}

func NewAirportRequest() *AirportRequest {
	return new(AirportRequest)
}

func (a *AirportRequest) AirportRequestToEntity(req *model.CreateAirportRequest) *entity.Airport {
	var entity entity.Airport
	if err := copier.Copy(&entity, &req); err != nil {
		panic(any(err))
	}

	return &entity
}

func (a *AirportRequest) GetAirportRequestToEntity(req *model.GetAirportRequest) *entity.Airport {
	var entity entity.Airport
	if err := copier.Copy(&entity, &req); err != nil {
		panic(any(err))
	}

	return &entity
}
