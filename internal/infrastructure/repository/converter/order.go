package converter

import (
	"flight_system/internal/domain/entity"
	"flight_system/internal/infrastructure/repository/po"
)

func OrderEntityToPo(order *entity.Order) *po.Order {
	return &po.Order{
		ID:        order.ID,
		TicketID:  order.TicketID,
		UserID:    order.UserID,
		Price:     order.Price,
		CreatedAt: order.CreatedAt,
	}
}

func OrderPoToEntity(order *po.Order) *entity.Order {
	return &entity.Order{
		ID:        order.ID,
		TicketID:  order.TicketID,
		UserID:    order.UserID,
		Price:     order.Price,
		CreatedAt: order.CreatedAt,
	}
}
