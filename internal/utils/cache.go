package utils

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/redis/go-redis/v9"
)

type (
	FN[T any] func() (data T, err error)

	// Cache Interface
	// maybe not just for redis
	Cache[T any] interface {
		CacheFirstData(ctx context.Context, rdp *redis.Client, i *CacheDataSupplied[T]) (data T, err error)
	}

	CacheDataSupplied[T any] struct {
		Key string
		TTL time.Duration
		CbF FN[T]
	}
)

// CacheFirstData tries to get data from Redis; if not found, it loads from the callback and stores it.
func CacheFirstData[T any](ctx context.Context, rdp *redis.Client, i *CacheDataSupplied[T]) (T, error) {
	var zero T
	// Try to load data from Redis
	valueCache, errCache := rdp.Get(ctx, i.Key).Result()
	// if error, load data from repository
	if errCache != nil {
		// load data from repository
		data, err := i.CbF()
		if err != nil {
			return zero, err
		}
		// if redis is connected and data is null save data from repository
		if errors.Is(errCache, redis.Nil) {
			// encode given data
			if jsonData, e := json.Marshal(data); e == nil {
				// store data to redis
				_ = rdp.Set(ctx, i.Key, jsonData, i.TTL).Err()
			}
		}
		// return back data from repository
		return data, nil
	}
	// Decode the cached data
	var result T
	if err := json.Unmarshal([]byte(valueCache), &result); err != nil {
		return zero, err
	}
	return result, nil
}
