package config

import (
	"log"

	"github.com/redis/go-redis/v9"
)

func RedisConnection() Option {
	return func(cfg *Config) {
		redisSingleton.Do(func() {
			log.Println("Trying to open redis connection pool . . . .")
			RedisPool = redis.NewClient(&redis.Options{
				Addr:     cfg.RedisDsnURL,
				Password: "",
				DB:       0,
			})
			if err := RedisPool.Ping(cfg.ctx).Err(); err != nil {
				log.Fatalf("REDIS_ERROR: %s\n",
					err.Error())
			}
			log.Println("Redis connection pool created . . . .")
		})
	}
}
