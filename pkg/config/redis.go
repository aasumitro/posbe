package config

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
	"log"
)

func (cfg Config) InitRedisConn() {
	log.Println("Trying to open redis connection . . . .")

	redisCacheOnce.Do(func() {
		RedisCache = redis.NewClient(&redis.Options{
			Addr:     cfg.CacheDsnUrl,
			Password: "", // no password set
			DB:       0,  // use default DB
		})

		if err := RedisCache.Ping(context.Background()).Err(); err != nil {
			panic(fmt.Sprintf("REDIS_ERROR: %s", err.Error()))
		}

		log.Printf("Cache connected with %s driver . . . .", cfg.CacheDriver)
	})
}
