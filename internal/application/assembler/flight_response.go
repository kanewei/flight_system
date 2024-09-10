package assembler

import (
	"flight_system/internal/application/model"
	"flight_system/internal/domain/entity"
)

type FlightResponse struct{}

func NewFlightResponse() *FlightResponse {
	return new(FlightResponse)
}

func flightToResponse(flight *entity.Flight) *model.GetFlightResponse {
	return &model.GetFlightResponse{
		ID:                     flight.ID,
		Departure:              flight.Departure,
		Arrival:                flight.Arrival,
		DepartureTime:          flight.DepartureTime,
		ArrivalTime:            flight.ArrivalTime,
		Price:                  flight.Price,
		AirportCode:            flight.AirportCode,
		Airport:                *airportToResponse(&flight.Airport),
		AirplaneID:             flight.AirplaneID,
		Airplane:               *airplaneToResponse(&flight.Airplane),
		Seats:                  flightSeatsToResponse(flight.Seats),
		RemainingOverSoldSeats: isRemainingOverSoldSeats(flight.Seats),
	}
}

func flightSeatsToResponse(seats []entity.FlightSeat) []model.GetFlightSeatsResponse {
	var response []model.GetFlightSeatsResponse
	for _, seat := range seats {
		response = append(response, model.GetFlightSeatsResponse{
			FlightID:   seat.FlightID,
			ID:         seat.SeatID,
			Class:      seat.Class,
			IsOverSold: seat.IsOverSold,
			Available:  seat.Available,
		})
	}
	return response
}

func isRemainingOverSoldSeats(seats []entity.FlightSeat) bool {
	for _, seat := range seats {
		if seat.Available && !seat.IsOverSold {
			return true
		}
	}
	return false
}

func (f *FlightResponse) CreateFlightToResponse(entity *entity.Flight) *model.CreateFlightResponse {
	return &model.CreateFlightResponse{
		ID: entity.ID,
	}
}

func (f *FlightResponse) GetFlightByIdToResponse(entity *entity.Flight) *model.GetFlightResponse {
	return flightToResponse(entity)
}

func (f *FlightResponse) SearchFlightToResponse(entities []*entity.Flight) []*model.GetFlightResponse {
	var flights []*model.GetFlightResponse
	for _, entity := range entities {
		flights = append(flights, flightToResponse(entity))
	}
	return flights
}
