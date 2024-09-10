package repository

import (
	"context"
	"flight_system/internal/domain/entity"
)

type TicketRepository interface {
	CreateTicket(ticket *entity.Ticket) error
	GetTicketById(id string) (*entity.Ticket, error)
	GetTicketByUserId(userId int64) ([]*entity.Ticket, error)
}

type TicketCacheRepository interface {
	CreateTicketOrder(ctx context.Context, ticketOrder *entity.TicketOrder) (string, error)
	GetUserHasTicketOrderById(ctx context.Context, userId int64, ticketId string) (*entity.Ticket, error)
	GetTicketsByUserId(ctx context.Context, userId int64) ([]*entity.Ticket, error)
}
