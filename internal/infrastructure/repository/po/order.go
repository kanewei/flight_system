package po

import "time"

type Order struct {
	ID        int64     `gorm:"primaryKey;column:id"`
	TicketID  string    `gorm:"ticket_id"`
	Ticket    Ticket    `gorm:"foreignKey:ticket_id"`
	UserID    int64     `gorm:"user_id"`
	User      User      `gorm:"foreignKey:user_id"`
	Price     string    `gorm:"price"`
	CreatedAt time.Time `gorm:"created_at"`
}
