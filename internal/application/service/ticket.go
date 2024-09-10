package service

import (
	"context"
	"flight_system/internal/application/assembler"
	"flight_system/internal/application/model"
	"flight_system/internal/domain/aggregate"
)

type TicketService struct {
	TicketRequest  *assembler.TicketRequest
	TicketResponse *assembler.TicketResponse
	TicketAggrate  aggregate.TicketAggrate
}

func NewTicketService() *TicketService {
	return &TicketService{
		TicketRequest:  assembler.NewTicketRequest(),
		TicketResponse: assembler.NewTicketResponse(),
		TicketAggrate:  aggregate.NewTicketAggrate(),
	}
}

func (t *TicketService) CreateTicketOrder(ctx context.Context, req *model.CreateTicketOrderRequest) (*model.CreateTicketOrderResponse, error) {
	ticketOrder := t.TicketRequest.CreateTicketOrderRequestToEntity(req)
	orderedTicket, err := t.TicketAggrate.CreateTicketOrder(ctx, ticketOrder)
	if err != nil {
		return nil, err
	}
	return t.TicketResponse.CreateTicketOrderResponseToEntity(orderedTicket), nil
}

func (t *TicketService) GetById(req *model.GetTicketRequest) (*model.GetTicketResponse, error) {
	ticket := t.TicketRequest.GetTicketRequestToEntity(req)
	ticket, err := t.TicketAggrate.GetTicketById(ticket.ID)
	if err != nil {
		return nil, err
	}
	return t.TicketResponse.GetTicketToResponse(ticket), nil
}

func (t *TicketService) GetByUserId(ctx context.Context, req *model.GetUserTicketRequest) ([]*model.GetUserTicketResponse, error) {
	ticket := t.TicketRequest.GetUserTicketRequestToEntity(req)
	tickets, err := t.TicketAggrate.GetTicketByUserId(ctx, ticket.UserID)
	if err != nil {
		return nil, err
	}
	return t.TicketResponse.GetUserTicketsToResponse(tickets), nil
}
