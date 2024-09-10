package service

import (
	"flight_system/internal/application/assembler"
	"flight_system/internal/application/model"
	"flight_system/internal/domain/aggregate"
)

type AirplaneService struct {
	airplaneRequest  *assembler.AirplaneRequest
	airplaneResponse *assembler.AirplaneResponse
	airplaneAggrate  aggregate.AirplaneAggrate
}

func NewAirplaneService() *AirplaneService {
	return &AirplaneService{
		airplaneRequest:  assembler.NewAirplaneRequest(),
		airplaneResponse: assembler.NewAirplaneResponse(),
		airplaneAggrate:  aggregate.NewAirplaneAggrate(),
	}
}

func (a *AirplaneService) Create(req *model.CreateAirplaneRequest) (*model.CreateAirplaneResponse, error) {
	airplane := a.airplaneRequest.AirplaneRequestToEntity(req)
	id, err := a.airplaneAggrate.Create(airplane)
	if err != nil {
		return nil, err
	}
	return a.airplaneResponse.CreateToResponse(id), nil
}

func (a *AirplaneService) GetAirplaneById(req *model.GetAirplaneRequest) (*model.GetAirplaneResponse, error) {
	airplane := a.airplaneRequest.GetAirplaneRequestToEntity(req)
	airplane, err := a.airplaneAggrate.GetAirplaneById(airplane.ID)
	if err != nil {
		return nil, err
	}
	return a.airplaneResponse.GetByIdToResponse(airplane), nil
}

func (a *AirplaneService) GetAirplaneByModel(req *model.GetAirplaneRequest) ([]*model.GetAirplaneResponse, error) {
	airplane := a.airplaneRequest.GetAirplaneRequestToEntity(req)
	airplanes, err := a.airplaneAggrate.GetAirplaneByModel(airplane.Model)
	if err != nil {
		return nil, err
	}
	return a.airplaneResponse.GetByModelToResponse(airplanes), nil
}
