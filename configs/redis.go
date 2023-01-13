package configs

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
	"log"
)

func (cfg Config) InitRedisConn() {
	log.Println("Trying to open redis connection pool . . . .")

	redisCacheOnce.Do(func() {
		RedisPool = redis.NewClient(&redis.Options{
			Addr:     cfg.CacheDsnURL,
			Password: "", // no password set
			DB:       0,  // use default DB
		})

		if err := RedisPool.Ping(context.Background()).Err(); err != nil {
			panic(fmt.Sprintf("REDIS_ERROR: %s", err.Error()))
		}

		log.Print("Redis connection pool created . . . .")
	})
}
