package assembler

import (
	"flight_system/internal/application/model"
	"flight_system/internal/domain/entity"
)

type OrderRequest struct{}

func NewOrderRequest() *OrderRequest {
	return new(OrderRequest)
}

func (o *OrderRequest) OrderRequestToEntity(req *model.CreateOrderRequest) *entity.Order {
	return &entity.Order{
		TicketID: req.TicketID,
	}
}
