package converter

import (
	"flight_system/internal/domain/entity"
	"flight_system/internal/infrastructure/repository/po"
)

func CreateFlightToPo(flight *entity.Flight) *po.Flight {
	return &po.Flight{
		Departure:     flight.Departure,
		Arrival:       flight.Arrival,
		DepartureTime: flight.DepartureTime,
		ArrivalTime:   flight.ArrivalTime,
		Price:         flight.Price,
		AirportCode:   flight.AirportCode,
		AirplaneID:    flight.AirplaneID,
	}
}

func CreateFlightSeatsToPo(flightId string, seats []*entity.FlightSeat) []*po.FlightSeat {
	var flightSeats []*po.FlightSeat
	for _, seat := range seats {
		flightSeats = append(flightSeats, &po.FlightSeat{
			FlightID:   flightId,
			SeatID:     seat.SeatID,
			Class:      seat.Class,
			IsOverSold: seat.IsOverSold,
			Available:  seat.Available,
		})
	}
	return flightSeats
}

func FlightPoToEntity(flight *po.Flight) (*entity.Flight, error) {
	airplane, err := AirplanePoToEntity(&flight.Airplane)
	if err != nil {
		return nil, err
	}

	return &entity.Flight{
		ID:            flight.ID,
		Departure:     flight.Departure,
		Arrival:       flight.Arrival,
		DepartureTime: flight.DepartureTime,
		ArrivalTime:   flight.ArrivalTime,
		Price:         flight.Price,
		AirportCode:   flight.AirportCode,
		Airport:       *AirportPoToEntity(&flight.Airport),
		AirplaneID:    flight.AirplaneID,
		Airplane:      *airplane,
		Seats:         FlightSeatsPoToEntities(flight.Seats),
	}, nil
}

func FlightsPoToEntities(flights []*po.Flight) ([]*entity.Flight, error) {
	entities := make([]*entity.Flight, 0, len(flights))
	for _, flight := range flights {
		entity, err := FlightPoToEntity(flight)
		if err != nil {
			return nil, err
		}
		entities = append(entities, entity)
	}
	return entities, nil
}

func FlightSeatsPoToEntities(flightSeats []po.FlightSeat) []entity.FlightSeat {
	entities := make([]entity.FlightSeat, 0, len(flightSeats))
	for _, seat := range flightSeats {
		entities = append(entities, entity.FlightSeat{
			FlightID:   seat.FlightID,
			SeatID:     seat.SeatID,
			Class:      seat.Class,
			IsOverSold: seat.IsOverSold,
			Available:  seat.Available,
		})
	}
	return entities
}
