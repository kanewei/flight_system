package po

type Ticket struct {
	ID         string     `gorm:"primaryKey;column:id"`
	UserID     int64      `gorm:"user_id"`
	User       User       `gorm:"foreignKey:user_id"`
	FlightID   string     `gorm:"flight_id"`
	Flight     Flight     `gorm:"foreignKey:flight_id"`
	SeatID     string     `gorm:"seat_id"`
	FlightSeat FlightSeat `gorm:"foreignKey:seat_id"`
	Price      string     `gorm:"price"`
	Status     int        `gorm:"status"`
}
