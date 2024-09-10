package database

import (
	"flight_system/internal/domain/entity"
	"flight_system/internal/global"
	"flight_system/internal/infrastructure/repository/converter"
	"flight_system/internal/infrastructure/repository/po"

	"gorm.io/gorm"
)

var pageSize = 10

type FlightRepo struct {
	flightTable string
	db          *gorm.DB
}

func NewFlightRepoImpl() *FlightRepo {
	return &FlightRepo{
		flightTable: "flights",
		db:          global.Db,
	}
}

func (f *FlightRepo) CreateFlight(flight *entity.Flight) (*entity.Flight, error) {
	po := converter.CreateFlightToPo(flight)
	result := f.db.Table(f.flightTable).Create(&po)
	if result.Error != nil {
		return nil, result.Error
	}
	return flight, nil
}

func (f *FlightRepo) GetFlightById(id string) (*entity.Flight, error) {
	var flight po.Flight
	result := f.db.Preload("Airport").Preload("Airplane").Preload("Seats").Find(&flight, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	flightEntity, err := converter.FlightPoToEntity(&flight)
	if err != nil {
		return nil, err
	}

	return flightEntity, nil
}

func (f *FlightRepo) GetFlightByIdWithAvaliableSeats(id string) (*entity.Flight, error) {
	var flight po.Flight
	result := f.db.Preload("Airport").Preload("Airplane").Preload("Seats", "avaliable = ?", true).Find(&flight, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	flightEntity, err := converter.FlightPoToEntity(&flight)
	if err != nil {
		return nil, err
	}

	return flightEntity, nil
}

func (f *FlightRepo) GetFlightByIdAndSeatId(flightId string, seatId string) (*entity.Flight, error) {
	var flightPo po.Flight
	result := f.db.Preload("Airport").Preload("Airplane").Preload("Seats", "id = ?", seatId).Find(&flightPo, "id = ?", flightId)
	if result.Error != nil {
		return nil, result.Error
	}

	flightEntity, err := converter.FlightPoToEntity(&flightPo)
	if err != nil {
		return nil, err
	}

	return flightEntity, nil
}

func (f *FlightRepo) SearchFlight(flight *entity.Flight, page int) ([]*entity.Flight, error) {
	var flights []*po.Flight
	result := f.db.Preload("Airport").Preload("Airplane").Preload("Seats").Limit(pageSize).Offset((page-1)*pageSize).Find(&flights, "departure = ? AND arrival = ? AND departure_time = ? AND arrival_time = ?",
		flight.Departure, flight.Arrival, flight.DepartureTime, flight.ArrivalTime)
	if result.Error != nil {
		return nil, result.Error
	}

	flightEntities, err := converter.FlightsPoToEntities(flights)
	if err != nil {
		return nil, err
	}

	return flightEntities, nil
}
