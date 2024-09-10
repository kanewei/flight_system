package service

import (
	"context"
	"flight_system/internal/application/assembler"
	"flight_system/internal/application/model"
	"flight_system/internal/domain/aggregate"
)

type OrderService struct {
	OrderRequest   *assembler.OrderRequest
	OrderResponse  *assembler.OrderResponse
	OrderAggregate aggregate.OrderAggregate
}

func NewOrderService() *OrderService {
	return &OrderService{
		OrderRequest:   assembler.NewOrderRequest(),
		OrderResponse:  assembler.NewOrderResponse(),
		OrderAggregate: aggregate.NewOrderAggregate(),
	}
}

func (o *OrderService) CreateOrder(ctx context.Context, req *model.CreateOrderRequest) (int64, error) {
	order := o.OrderRequest.OrderRequestToEntity(req)
	id, err := o.OrderAggregate.CreateOrder(ctx, order)
	if err != nil {
		return 0, err
	}
	return id, nil
}
