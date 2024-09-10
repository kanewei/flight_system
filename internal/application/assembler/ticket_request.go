package assembler

import (
	"flight_system/internal/application/model"
	"flight_system/internal/domain/entity"
)

type TicketRequest struct{}

func NewTicketRequest() *TicketRequest {
	return new(TicketRequest)
}

func (t *TicketRequest) CreateTicketOrderRequestToEntity(req *model.CreateTicketOrderRequest) *entity.TicketOrder {
	return &entity.TicketOrder{
		FlightID: req.FlightID,
		SeatID:   req.SeatID,
		UserID:   req.UserID,
	}
}

func (t *TicketRequest) CreateTicketRequestToEntity(req *model.CreateTicketRequest) *entity.TicketOrder {
	return &entity.TicketOrder{
		FlightID: req.FlightID,
		SeatID:   req.SeatID,
	}
}

func (t *TicketRequest) GetTicketRequestToEntity(req *model.GetTicketRequest) *entity.Ticket {
	return &entity.Ticket{
		ID: req.ID,
	}
}

func (t *TicketRequest) GetUserTicketRequestToEntity(req *model.GetUserTicketRequest) *entity.Ticket {
	return &entity.Ticket{
		UserID: req.UserID,
	}
}
