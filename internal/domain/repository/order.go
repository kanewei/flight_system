package repository

import "flight_system/internal/domain/entity"

type OrderRepository interface {
	CreateOrder(order *entity.Order) (int64, error)
}
