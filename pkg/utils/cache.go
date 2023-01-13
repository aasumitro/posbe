package utils

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v9"
	"time"
)

type (
	FNCache func() (data any, err error)

	// Cache Interface
	// maybe not just for redis
	Cache interface {
		CacheFirstData(i *CacheDataSupplied) (data any, err error)
	}

	RedisCache struct {
		Ctx     context.Context
		RdpConn *redis.Client
	}

	CacheDataSupplied struct {
		Key string
		TTL time.Duration
		CbF FNCache
	}
)

func (cache *RedisCache) CacheFirstData(i *CacheDataSupplied) (data any, err error) {
	// load data from redis
	valueCache, errCache := cache.RdpConn.Get(cache.Ctx, i.Key).Result()
	// if error, load data from repository
	if errCache != nil {
		// load data from repository
		data, err = i.CbF()
		// if redis is connected and data is null save data from repository
		if errCache == redis.Nil {
			// encode given data
			jsonData, _ := json.Marshal(data)
			// store data to redis
			cache.RdpConn.Set(cache.Ctx, i.Key, jsonData, i.TTL)
		}
		// return back data from repository
		return data, err
	}
	// return data
	return valueCache, nil
}
