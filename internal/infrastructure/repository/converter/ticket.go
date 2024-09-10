package converter

import (
	"flight_system/internal/domain/entity"
	"flight_system/internal/infrastructure/repository/po"
)

func TicketEntityToPo(entity *entity.Ticket) *po.Ticket {
	return &po.Ticket{
		ID:       entity.ID,
		FlightID: entity.FlightID,
		SeatID:   entity.SeatID,
		UserID:   entity.UserID,
		Price:    entity.Price,
		Status:   int(entity.Status),
	}
}

func TicketPoToEntity(po *po.Ticket) *entity.Ticket {
	return &entity.Ticket{
		ID:       po.ID,
		FlightID: po.FlightID,
		SeatID:   po.SeatID,
		UserID:   po.UserID,
		Price:    po.Price,
		Status:   entity.TicketStatus(po.Status),
	}
}

func TicketsToEntities(tickets []*po.Ticket) []*entity.Ticket {
	var entities []*entity.Ticket
	for _, ticket := range tickets {
		entities = append(entities, TicketPoToEntity(ticket))
	}
	return entities
}
