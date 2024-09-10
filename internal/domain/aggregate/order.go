package aggregate

import (
	"context"
	"flight_system/internal/domain/entity"
	"flight_system/internal/domain/repository"
	orderRepoImpl "flight_system/internal/infrastructure/repository/impl/order/database"
	ticketCacheRepoImpl "flight_system/internal/infrastructure/repository/impl/ticket/cache"
	ticketRepoImpl "flight_system/internal/infrastructure/repository/impl/ticket/database"
	"time"
)

type OrderAggregate interface {
	CreateOrder(ctx context.Context, order *entity.Order) (int64, error)
}

type orderAggregate struct {
	orderRepoImpl       repository.OrderRepository
	ticketRepoImpl      repository.TicketRepository
	ticketCacheRepoImpl repository.TicketCacheRepository
}

func NewOrderAggregate() *orderAggregate {
	return &orderAggregate{
		orderRepoImpl:       orderRepoImpl.NewOrderRepoImpl(),
		ticketRepoImpl:      ticketRepoImpl.NewTicketRepoImpl(),
		ticketCacheRepoImpl: ticketCacheRepoImpl.NewTicketCacheRepoImpl(),
	}
}

func (o *orderAggregate) CreateOrder(ctx context.Context, order *entity.Order) (int64, error) {
	ticket, err := o.ticketCacheRepoImpl.GetUserHasTicketOrderById(ctx, order.UserID, order.TicketID)
	if err != nil {
		return 0, err
	}

	// Todo payment process

	if ticket != nil {
		ticket.Status = entity.TicketStatusBooked
		err := o.ticketRepoImpl.CreateTicket(ticket)
		if err != nil {
			return 0, err
		}
	}

	createOrder := &entity.Order{
		TicketID:  order.TicketID,
		UserID:    order.UserID,
		Price:     ticket.Price,
		CreatedAt: time.Now(),
	}

	return o.orderRepoImpl.CreateOrder(createOrder)
}
