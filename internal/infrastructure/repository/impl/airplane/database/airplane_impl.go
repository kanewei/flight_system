package database

import (
	"flight_system/internal/domain/entity"
	"flight_system/internal/global"
	"flight_system/internal/infrastructure/repository/converter"
	"flight_system/internal/infrastructure/repository/po"

	"gorm.io/gorm"
)

type AirplaneRepo struct {
	AirplaneTable string
	db            *gorm.DB
}

func NewAirplaneRepoImpl() *AirplaneRepo {
	return &AirplaneRepo{
		AirplaneTable: "airplanes",
		db:            global.Db,
	}
}

func (a *AirplaneRepo) CreateAirplane(airplane *entity.Airplane) (int64, error) {
	po, err := converter.CreateAirplaneToPo(airplane)
	if err != nil {
		return 0, err
	}

	result := a.db.Table(a.AirplaneTable).Create(&po)
	if result.Error != nil {
		return 0, result.Error
	}
	return po.ID, nil
}

func (a *AirplaneRepo) GetAirplaneById(id int64) (*entity.Airplane, error) {
	var airplane po.Airplane
	result := a.db.Table(a.AirplaneTable).Where("id = ?", id).First(&airplane)
	if result.Error != nil {
		return nil, result.Error
	}

	airplaneEntity, err := converter.AirplanePoToEntity(&airplane)
	if err != nil {
		return nil, err
	}

	return airplaneEntity, nil
}

func (a *AirplaneRepo) GetAirplaneByModel(model string) ([]*entity.Airplane, error) {
	var airplanes []*po.Airplane
	result := a.db.Table(a.AirplaneTable).Where("model = ?", model).Find(&airplanes)
	if result.Error != nil {
		return nil, result.Error
	}

	airplaneEntities, err := converter.AirplanesPoToEntities(airplanes)
	if err != nil {
		return nil, err
	}
	return airplaneEntities, nil
}
