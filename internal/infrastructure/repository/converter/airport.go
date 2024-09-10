package converter

import (
	"flight_system/internal/domain/entity"
	"flight_system/internal/infrastructure/repository/po"
)

func CreateAirportToPo(airport *entity.Airport) *po.Airport {
	return &po.Airport{
		Code:      airport.Code,
		Terminal:  airport.Terminal,
		City:      airport.City,
		Name:      airport.Name,
		Available: airport.Available,
	}
}

func AirportPoToEntity(airport *po.Airport) *entity.Airport {
	return &entity.Airport{
		Code:      airport.Code,
		Terminal:  airport.Terminal,
		Name:      airport.Name,
		City:      airport.City,
		Available: airport.Available,
	}
}

func AirportsPoToEntities(airports []*po.Airport) []*entity.Airport {
	var entities []*entity.Airport
	for _, airport := range airports {
		entities = append(entities, AirportPoToEntity(airport))
	}
	return entities
}
