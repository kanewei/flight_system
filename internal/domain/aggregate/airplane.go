package aggregate

import (
	"flight_system/internal/domain/entity"
	"flight_system/internal/domain/repository"
	airplaneRepoImpl "flight_system/internal/infrastructure/repository/impl/airplane/database"
)

type AirplaneAggrate interface {
	Create(airplane *entity.Airplane) (int64, error)
	GetAirplaneById(id int64) (*entity.Airplane, error)
	GetAirplaneByModel(model string) ([]*entity.Airplane, error)
}

type airplaneAggrate struct {
	airplaneRepo repository.AirplaneRepository
}

func NewAirplaneAggrate() *airplaneAggrate {
	return &airplaneAggrate{
		airplaneRepo: airplaneRepoImpl.NewAirplaneRepoImpl(),
	}
}

func (a *airplaneAggrate) Create(airplane *entity.Airplane) (int64, error) {
	return a.airplaneRepo.CreateAirplane(airplane)
}

func (a *airplaneAggrate) GetAirplaneById(id int64) (*entity.Airplane, error) {
	return a.airplaneRepo.GetAirplaneById(id)
}

func (a *airplaneAggrate) GetAirplaneByModel(model string) ([]*entity.Airplane, error) {
	return a.airplaneRepo.GetAirplaneByModel(model)
}
