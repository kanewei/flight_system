package cache

import (
	"context"
	"flight_system/internal/domain/entity"
	"flight_system/internal/global"
	"fmt"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
)

type TicketCacheRepo struct {
	cacheDb *redis.Client
}

func NewTicketCacheRepoImpl() *TicketCacheRepo {
	return &TicketCacheRepo{
		cacheDb: global.RedisClient,
	}
}

func createTicketId(flightId, seatId string) string {
	return fmt.Sprintf("%v:%v", flightId, seatId)
}

func createTicketOrderKey(ticketId string) string {
	return fmt.Sprintf("ordered:%v", ticketId)
}

func createUserTicketKey(userId int64, ticketId string) string {
	return fmt.Sprintf("%v:%v", userId, ticketId)
}

func (t *TicketCacheRepo) CreateTicketOrder(ctx context.Context, ticketOrder *entity.TicketOrder) (string, error) {
	ticketId := createTicketId(ticketOrder.FlightID, ticketOrder.SeatID)
	exist, err := t.cacheDb.Exists(ctx, ticketId).Result()
	if err != nil {
		return "", err
	}

	if exist == 0 {
		return "", fmt.Errorf("flight ticket not found")
	}

	isAvailable, err := t.cacheDb.Get(ctx, ticketId).Bool()
	if err != nil {
		return "", err
	}

	// seat is available
	if !isAvailable {
		return "", fmt.Errorf("flight ticket is not available")
	}

	// check if ticket has been ordered
	ticketOrderedKey := createTicketOrderKey(ticketId)
	if exist, err := t.cacheDb.Exists(ctx, ticketOrderedKey).Result(); err != nil {
		return "", err
	} else if exist != 0 {
		// ticket has been ordered
		return "", fmt.Errorf("flight ticket has been ordered")
	}

	// create ticket order cache for 30 minutes
	if err := t.cacheDb.Set(ctx, ticketOrderedKey, ticketOrder.Price, 30*time.Minute).Err(); err != nil {
		return "", err
	}

	// create user ticket cache for 30 minutes
	userTicketKey := createUserTicketKey(ticketOrder.UserID, ticketId)
	if err := t.cacheDb.Set(ctx, userTicketKey, ticketId, 30*time.Minute).Err(); err != nil {
		return "", err
	}

	return ticketId, nil
}

func (t *TicketCacheRepo) GetUserHasTicketOrderById(ctx context.Context, userId int64, ticketId string) (*entity.Ticket, error) {
	userTicketKey := createUserTicketKey(userId, ticketId)
	exist, err := t.cacheDb.Exists(ctx, userTicketKey).Result()
	if err != nil {
		return nil, err
	}

	if exist == 0 {
		return nil, fmt.Errorf("ticket order not found")
	}

	tickerOrderKey := createTicketOrderKey(ticketId)
	price, err := t.cacheDb.Get(ctx, tickerOrderKey).Result()
	if err != nil {
		return nil, err
	}

	keyStruct := strings.Split(ticketId, ":")
	flightId := keyStruct[0]
	seatId := keyStruct[1]

	return &entity.Ticket{
		ID:       ticketId,
		FlightID: flightId,
		SeatID:   seatId,
		UserID:   userId,
		Price:    price,
	}, nil
}

func (t *TicketCacheRepo) GetTicketsByUserId(ctx context.Context, userId int64) ([]*entity.Ticket, error) {
	userTicketKey := fmt.Sprintf("%v:*", userId)
	keys, err := t.cacheDb.Keys(ctx, userTicketKey).Result()
	if err != nil {
		return nil, err
	}

	tickets := make([]*entity.Ticket, len(keys))
	values, err := t.cacheDb.MGet(ctx, keys...).Result()
	for i, key := range keys {
		price := values[i].(string)

		keyStruct := strings.Split(key, ":")
		flightId := keyStruct[1]
		seatId := keyStruct[2]

		tickets[i] = &entity.Ticket{
			ID:       key,
			FlightID: flightId,
			SeatID:   seatId,
			UserID:   userId,
			Price:    price,
			Status:   entity.TicketStatusOrdered,
		}

	}

	return tickets, nil
}
