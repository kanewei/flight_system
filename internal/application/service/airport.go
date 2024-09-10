package service

import (
	"flight_system/internal/application/assembler"
	"flight_system/internal/application/model"
	"flight_system/internal/domain/aggregate"
)

type AirportService struct {
	airPortRequest  *assembler.AirportRequest
	airPortResponse *assembler.AirportResponse
	airPortAggrate  aggregate.AirportAggrate
}

func NewAirportService() *AirportService {
	return &AirportService{
		airPortRequest:  assembler.NewAirportRequest(),
		airPortResponse: assembler.NewAirportResponse(),
		airPortAggrate:  aggregate.NewAirportAggrate(),
	}
}

func (a *AirportService) CreateAirport(req *model.CreateAirportRequest) error {
	airport := a.airPortRequest.AirportRequestToEntity(req)
	err := a.airPortAggrate.CreateAirport(airport)
	if err != nil {
		return err
	}
	return nil
}

func (a *AirportService) GetAirportByCode(req *model.GetAirportRequest) (*model.GetAirportResponse, error) {
	airport := a.airPortRequest.GetAirportRequestToEntity(req)
	airport, err := a.airPortAggrate.GetAirportByCode(airport.Code)
	if err != nil {
		return nil, err
	}
	return a.airPortResponse.GetByCodeToResponse(airport), nil
}

func (a *AirportService) GetAirportByCity(req *model.GetAirportRequest) ([]*model.GetAirportResponse, error) {
	airport := a.airPortRequest.GetAirportRequestToEntity(req)
	airports, err := a.airPortAggrate.GetAirportByCity(airport.City)
	if err != nil {
		return nil, err
	}
	return a.airPortResponse.GetByCityToResponse(airports), nil
}
