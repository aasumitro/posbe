package config

import (
	"github.com/go-redis/redis/v9"
	"log"
)

var redisCache *redis.Client

func (cfg Config) InitRedisConn() {
	log.Println("Trying to open redis connection . . . .")

	if cfg.CacheDriver == "redis" {
		redisCache = redis.NewClient(&redis.Options{
			Addr:     cfg.CacheDsnUrl,
			Password: "", // no password set
			DB:       0,  // use default DB
		})
	}

	log.Printf("Cache connected with %s driver . . . .", cfg.CacheDriver)
}

func (cfg Config) GetRedisConnection() *redis.Client {
	if cfg.CacheDriver == "redis" {
		return redisCache
	}

	return nil
}
