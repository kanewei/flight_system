package converter

import (
	"encoding/json"
	"flight_system/internal/domain/entity"
	"flight_system/internal/infrastructure/repository/po"
)

func CreateAirplaneToPo(airplane *entity.Airplane) (*po.Airplane, error) {
	seats, err := json.Marshal(airplane.Seats)
	if err != nil {
		return nil, err
	}
	return &po.Airplane{
		Model: airplane.Model,
		Seats: seats,
	}, nil
}

func AirplanePoToEntity(airplane *po.Airplane) (*entity.Airplane, error) {
	var seats []*entity.Seat
	if err := json.Unmarshal(airplane.Seats, &seats); err != nil {
		return nil, err
	}
	return &entity.Airplane{
		ID:    airplane.ID,
		Model: airplane.Model,
		Seats: seats,
	}, nil
}

func AirplanesPoToEntities(airplanes []*po.Airplane) ([]*entity.Airplane, error) {
	entities := make([]*entity.Airplane, 0, len(airplanes))
	for _, airplane := range airplanes {
		entity, err := AirplanePoToEntity(airplane)
		if err != nil {
			return nil, err
		}
		entities = append(entities, entity)
	}
	return entities, nil
}
