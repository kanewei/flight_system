package database

import (
	"flight_system/internal/domain/entity"
	"flight_system/internal/global"
	"flight_system/internal/infrastructure/repository/converter"

	"gorm.io/gorm"
)

type FlightSeatRepoImpl struct {
	flightSeatTable string
	db              *gorm.DB
}

func NewFlightSeatRepoImpl() *FlightSeatRepoImpl {
	return &FlightSeatRepoImpl{
		flightSeatTable: "flight_seats",
		db:              global.Db,
	}
}

func (f *FlightSeatRepoImpl) CreateFlightSeats(flightId string, flightSeats []*entity.FlightSeat) error {
	po := converter.CreateFlightSeatsToPo(flightId, flightSeats)
	result := f.db.Table(f.flightSeatTable).Create(&po)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
