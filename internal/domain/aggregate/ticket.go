package aggregate

import (
	"context"
	"flight_system/internal/domain/entity"
	"flight_system/internal/domain/repository"
	flightRepoImpl "flight_system/internal/infrastructure/repository/impl/flight/database"
	ticketCacheRepoImpl "flight_system/internal/infrastructure/repository/impl/ticket/cache"
	ticketRepoImpl "flight_system/internal/infrastructure/repository/impl/ticket/database"
	"flight_system/pkg/calculator"

	"math/rand"
)

type TicketAggrate interface {
	CreateTicketOrder(ctx context.Context, ticketOrder *entity.TicketOrder) (*entity.TicketOrder, error)
	CreateTicket(ticket *entity.Ticket) error
	GetTicketById(id string) (*entity.Ticket, error)
	GetTicketByUserId(ctx context.Context, userId int64) ([]*entity.Ticket, error)
}

type ticketAggrate struct {
	ticketRepoImpl      repository.TicketRepository
	ticketCacheRepoImpl repository.TicketCacheRepository
	flightRepoImpl      repository.FlightRepository
}

func NewTicketAggrate() *ticketAggrate {
	return &ticketAggrate{
		ticketRepoImpl:      ticketRepoImpl.NewTicketRepoImpl(),
		ticketCacheRepoImpl: ticketCacheRepoImpl.NewTicketCacheRepoImpl(),
		flightRepoImpl:      flightRepoImpl.NewFlightRepoImpl(),
	}
}

func (t *ticketAggrate) CreateTicketOrder(ctx context.Context, ticketOrder *entity.TicketOrder) (ticketOrdered *entity.TicketOrder, err error) {
	var flight *entity.Flight
	var seat entity.FlightSeat
	times := 1.0

	// if seatID is empty, get random seat and is possible to get seat is over sold
	if ticketOrder.SeatID == "" {
		flight, err = t.flightRepoImpl.GetFlightByIdWithAvaliableSeats(ticketOrder.FlightID)
		if err != nil {
			return nil, err
		}
		seat = getRandomSeat(flight.Seats)
		ticketOrder.SeatID = seat.SeatID
		times = 0.7
	} else {
		flight, err = t.flightRepoImpl.GetFlightByIdAndSeatId(ticketOrder.FlightID, ticketOrder.SeatID)
		if err != nil {
			return nil, err
		}
		seat = flight.Seats[0]
	}

	switch seat.Class {
	case "first":
		times = 2.0
	case "business":
		times = 1.5
	case "economy":
		times = 1.0
	}

	price, err := calculator.MutipleString(ticketOrder.Price, times)
	if err != nil {
		return nil, err
	}
	ticketOrdered = &entity.TicketOrder{
		FlightID: ticketOrder.FlightID,
		SeatID:   ticketOrder.SeatID,
		UserID:   ticketOrder.UserID,
	}
	ticketOrdered.Price = price
	ticketOrder.Price = price

	ticketId, err := t.ticketCacheRepoImpl.CreateTicketOrder(ctx, ticketOrder)
	if err != nil {
		return nil, err
	}

	ticketOrdered.ID = ticketId

	return ticketOrdered, nil
}

func getRandomSeat(seats []entity.FlightSeat) entity.FlightSeat {
	return seats[rand.Intn(len(seats))]
}

func (t *ticketAggrate) CreateTicket(ticket *entity.Ticket) error {
	return t.ticketRepoImpl.CreateTicket(ticket)
}

func (t *ticketAggrate) GetTicketById(id string) (*entity.Ticket, error) {
	return t.ticketRepoImpl.GetTicketById(id)
}

func (t *ticketAggrate) GetTicketByUserId(ctx context.Context, userId int64) ([]*entity.Ticket, error) {
	tickets, err := t.ticketRepoImpl.GetTicketByUserId(userId)
	if err != nil {
		return nil, err
	}

	orderedTickets, err := t.ticketCacheRepoImpl.GetTicketsByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}
	tickets = append(tickets, orderedTickets...)

	return tickets, nil
}
