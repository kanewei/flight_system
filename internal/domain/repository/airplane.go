package repository

import (
	"flight_system/internal/domain/entity"
)

type AirplaneRepository interface {
	CreateAirplane(airplane *entity.Airplane) (int64, error)
	GetAirplaneById(id int64) (*entity.Airplane, error)
	GetAirplaneByModel(model string) ([]*entity.Airplane, error)
}
