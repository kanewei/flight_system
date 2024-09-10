package database

import (
	"flight_system/internal/domain/entity"
	"flight_system/internal/global"
	"flight_system/internal/infrastructure/repository/converter"
	"flight_system/internal/infrastructure/repository/po"

	"gorm.io/gorm"
)

type TicketRepo struct {
	TicketTable string
	db          *gorm.DB
}

func NewTicketRepoImpl() *TicketRepo {
	return &TicketRepo{
		TicketTable: "tickets",
		db:          global.Db,
	}
}

func (t *TicketRepo) CreateTicket(ticket *entity.Ticket) error {
	po := converter.TicketEntityToPo(ticket)
	result := t.db.Table(t.TicketTable).Create(&po)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (t *TicketRepo) GetTicketById(id string) (*entity.Ticket, error) {
	var ticket po.Ticket
	result := t.db.Preload("User").Preload("Flight").Preload("FlightSeat").Find(&ticket, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return converter.TicketPoToEntity(&ticket), nil
}

func (t *TicketRepo) GetTicketByUserId(userId int64) ([]*entity.Ticket, error) {
	var tickets []*po.Ticket
	result := t.db.Preload("User").Preload("Flight").Preload("FlightSeat").Find(&tickets, "user_id = ?", userId)
	if result.Error != nil {
		return nil, result.Error
	}
	return converter.TicketsToEntities(tickets), nil
}
