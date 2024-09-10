package database

import (
	"flight_system/internal/domain/entity"
	"flight_system/internal/global"
	"flight_system/internal/infrastructure/repository/converter"

	"gorm.io/gorm"
)

type OrderRepo struct {
	OrderTable string
	db         *gorm.DB
}

func NewOrderRepoImpl() *OrderRepo {
	return &OrderRepo{
		OrderTable: "orders",
		db:         global.Db,
	}
}

func (o *OrderRepo) CreateOrder(order *entity.Order) (int64, error) {
	po := converter.OrderEntityToPo(order)
	result := o.db.Table(o.OrderTable).Create(&po)
	if result.Error != nil {
		return 0, result.Error
	}
	return po.ID, nil
}
