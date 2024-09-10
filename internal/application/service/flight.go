package service

import (
	"context"
	"flight_system/internal/application/assembler"
	"flight_system/internal/application/model"
	"flight_system/internal/domain/aggregate"
)

type FlightService struct {
	flightRequest    *assembler.FlightRequest
	flightRespone    *assembler.FlightResponse
	flightAggAggrate aggregate.FlightAggrate
}

func NewFlightService() *FlightService {
	return &FlightService{
		flightRequest:    assembler.NewFlightRequest(),
		flightRespone:    assembler.NewFlightResponse(),
		flightAggAggrate: aggregate.NewFlightAggrate(),
	}
}

func (f *FlightService) CreateFlight(ctx context.Context, req *model.CreateFlightRequest) (string, error) {
	flight := f.flightRequest.FlightRequestToEntity(req)
	id, err := f.flightAggAggrate.CreateFlight(ctx, flight)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (f *FlightService) GetFlightById(req *model.GetFlightRequest) (*model.GetFlightResponse, error) {
	flight := f.flightRequest.GetFlightRequestToEntity(req)
	flight, err := f.flightAggAggrate.GetFlightById(flight.ID)
	if err != nil {
		return nil, err
	}
	return f.flightRespone.GetFlightByIdToResponse(flight), nil
}

func (f *FlightService) SearchFlight(req *model.SearchFlightRequest) ([]*model.GetFlightResponse, error) {
	flight := f.flightRequest.SearchFlightRequestToEntity(req)
	flights, err := f.flightAggAggrate.SearchFlight(flight, req.Page)
	if err != nil {
		return nil, err
	}
	return f.flightRespone.SearchFlightToResponse(flights), nil
}
