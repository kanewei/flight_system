package po

import "time"

type Flight struct {
	ID            string       `gorm:"primaryKey;column:id"`
	Departure     string       `gorm:"departure"`
	Arrival       string       `gorm:"arrival"`
	DepartureTime time.Time    `gorm:"departure_time"`
	ArrivalTime   time.Time    `gorm:"arrival_time"`
	Price         string       `gorm:"price"`
	AirportCode   string       `gorm:"airport_code"`
	Airport       Airport      `gorm:"foreignKey:airport_code"`
	AirplaneID    int64        `gorm:"airplane_id"`
	Airplane      Airplane     `gorm:"foreignKey:airplane_id"`
	Seats         []FlightSeat `gorm:"foreignKey:flight_id"`
}

type FlightSeat struct {
	FlightID   string `gorm:"flight_id"`
	SeatID     string `gorm:"seat_id"`
	Class      string `gorm:"class"`
	IsOverSold bool   `gorm:"is_over_sold"`
	Available  bool   `gorm:"available"`
}
