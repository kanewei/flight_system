package assembler

import (
	"flight_system/internal/application/model"
	"flight_system/internal/domain/entity"
)

type OrderResponse struct{}

func NewOrderResponse() *OrderResponse {
	return new(OrderResponse)
}

func (c *OrderResponse) OrderToResponse(entity *entity.Order) *model.CreateOrderResponse {
	return &model.CreateOrderResponse{
		ID: entity.ID,
	}
}
