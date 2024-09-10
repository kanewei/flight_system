package cache

import (
	"context"
	"flight_system/internal/global"
	"time"

	"github.com/go-redis/redis/v8"
)

type FlightTicketCacheRepository struct {
	cacheDb *redis.Client
}

func NewFlightTicketCacheRepoImpl() *FlightTicketCacheRepository {
	return &FlightTicketCacheRepository{
		cacheDb: global.RedisClient,
	}
}

func (f *FlightTicketCacheRepository) CreateFlightTickets(ctx context.Context, flightTickets map[string]interface{}, expireTime time.Duration) error {
	_, err := f.cacheDb.MSet(ctx, flightTickets).Result()
	if err != nil {
		return err
	}

	for ticketId := range flightTickets {
		err = f.cacheDb.Expire(ctx, ticketId, expireTime).Err()
		if err != nil {
			return err
		}
	}
	return nil
}
