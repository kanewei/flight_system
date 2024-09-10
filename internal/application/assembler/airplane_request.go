package assembler

import (
	"flight_system/internal/application/model"
	"flight_system/internal/domain/entity"

	"github.com/jinzhu/copier"
)

type AirplaneRequest struct{}

func NewAirplaneRequest() *AirplaneRequest {
	return new(AirplaneRequest)
}

func (a *AirplaneRequest) AirplaneRequestToEntity(req *model.CreateAirplaneRequest) *entity.Airplane {
	var entity entity.Airplane
	if err := copier.Copy(&entity, &req); err != nil {
		panic(any(err))
	}

	return &entity
}

func (a *AirplaneRequest) GetAirplaneRequestToEntity(req *model.GetAirplaneRequest) *entity.Airplane {
	var entity entity.Airplane
	if err := copier.Copy(&entity, &req); err != nil {
		panic(any(err))
	}

	return &entity
}
