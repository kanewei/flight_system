package cache

import (
	"context"
	"time"

	"encoding/json"
	"flight_system/internal/domain/entity"
	"flight_system/internal/global"
	"flight_system/pkg/random"
	"strings"

	"github.com/go-redis/redis/v8"
)

var pageSize = 10

type FlightCacheRepository struct {
	cacheDb *redis.Client
}

func NewFlightCacheRepoImpl() *FlightCacheRepository {
	return &FlightCacheRepository{
		cacheDb: global.RedisClient,
	}
}

func createSearchFlightKey(flight *entity.Flight) string {
	var builder strings.Builder
	builder.WriteString(flight.Departure)
	builder.WriteString(":")
	builder.WriteString(flight.Arrival)
	builder.WriteString(":")
	builder.WriteString(flight.DepartureTime.Format("2006-01-02"))
	builder.WriteString(":")
	builder.WriteString(flight.ArrivalTime.Format("2006-01-02"))

	return builder.String()
}

func (f *FlightCacheRepository) GetSearchFlight(ctx context.Context, flight *entity.Flight, page int) ([]*entity.Flight, error) {
	return f.getPaginatedFlights(ctx, flight, page), nil
}

func (f *FlightCacheRepository) SetSearchFlight(ctx context.Context, flight *entity.Flight, flights []*entity.Flight) error {
	key := createSearchFlightKey(flight)

	cachedFlights := []*redis.Z{}
	for i, flight := range flights {
		member, err := json.Marshal(flight)
		if err != nil {
			return err
		}
		cachedFlights = append(cachedFlights, &redis.Z{Score: float64(i), Member: string(member)})
	}
	f.cacheDb.ZAdd(ctx, key, cachedFlights...)

	randomNumber := random.RandRange(10, 30)
	f.cacheDb.Expire(ctx, key, time.Duration(randomNumber)*time.Minute)

	return nil
}

func (f *FlightCacheRepository) getPaginatedFlights(ctx context.Context, flight *entity.Flight, page int) []*entity.Flight {
	key := createSearchFlightKey(flight)

	start := (page - 1) * pageSize
	end := start + pageSize - 1
	items, _ := f.cacheDb.ZRange(ctx, key, int64(start), int64(end)).Result()

	flights := make([]*entity.Flight, len(items))
	for i, item := range items {
		err := json.Unmarshal([]byte(item), &flights[i])
		if err != nil {
			return nil
		}
	}
	return flights
}
