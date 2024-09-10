package database

import (
	"flight_system/internal/domain/entity"
	"flight_system/internal/global"
	"flight_system/internal/infrastructure/repository/converter"
	"flight_system/internal/infrastructure/repository/po"

	"gorm.io/gorm"
)

type AirportRepo struct {
	AirportTable string
	db           *gorm.DB
}

func NewAirportRepoImpl() *AirportRepo {
	return &AirportRepo{
		AirportTable: "airports",
		db:           global.Db,
	}
}

func (a *AirportRepo) CreateAirport(airport *entity.Airport) error {
	po := converter.CreateAirportToPo(airport)
	result := a.db.Table(a.AirportTable).Create(&po)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (a *AirportRepo) GetAirportByCode(code string) (*entity.Airport, error) {
	var airport po.Airport
	result := a.db.Table(a.AirportTable).Where("code = ?", code).First(&airport)
	if result.Error != nil {
		return nil, result.Error
	}

	return converter.AirportPoToEntity(&airport), nil
}

func (a *AirportRepo) GetAirportByCity(city string) ([]*entity.Airport, error) {
	var airports []*po.Airport
	result := a.db.Table(a.AirportTable).Where("city = ?", city).Find(airports)
	if result.Error != nil {
		return nil, result.Error
	}
	return converter.AirportsPoToEntities(airports), nil
}
