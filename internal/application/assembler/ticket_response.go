package assembler

import (
	"flight_system/internal/application/model"
	"flight_system/internal/domain/entity"
)

type TicketResponse struct{}

func NewTicketResponse() *TicketResponse {
	return new(TicketResponse)
}

func (t *TicketResponse) CreateTicketOrderResponseToEntity(ticketOrder *entity.TicketOrder) *model.CreateTicketOrderResponse {
	return &model.CreateTicketOrderResponse{
		ID:    ticketOrder.ID,
		Price: ticketOrder.Price,
	}
}

func (t *TicketResponse) CreateTicketToResponse(ticketID string) *model.CreateTicketResponse {
	return &model.CreateTicketResponse{
		ID: ticketID,
	}
}

func (t *TicketResponse) GetTicketToResponse(ticket *entity.Ticket) *model.GetTicketResponse {
	return &model.GetTicketResponse{
		ID:     ticket.ID,
		Flight: *flightToResponse(&ticket.Flight),
	}
}

func (t *TicketResponse) GetTicketsToResponse(tickets []*entity.Ticket) []*model.GetTicketResponse {
	var response []*model.GetTicketResponse
	for _, ticket := range tickets {
		response = append(response, t.GetTicketToResponse(ticket))
	}
	return response
}

func (t *TicketResponse) GetUserTicketToResponse(ticket *entity.Ticket) *model.GetUserTicketResponse {
	return &model.GetUserTicketResponse{
		ID:     ticket.ID,
		Flight: *flightToResponse(&ticket.Flight),
	}
}

func (t *TicketResponse) GetUserTicketsToResponse(tickets []*entity.Ticket) []*model.GetUserTicketResponse {
	var response []*model.GetUserTicketResponse
	for _, ticket := range tickets {
		response = append(response, t.GetUserTicketToResponse(ticket))
	}
	return response
}
