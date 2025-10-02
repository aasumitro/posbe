package config

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

func PostgresConnection() Option {
	return func(cfg *Config) {
		conn, err := pgxpool.New(cfg.ctx, cfg.PostgresDsnURL)
		if err != nil {
			log.Fatalf("POSTGRES_ERROR: %s\n", err.Error())
		}
		if err := conn.Ping(cfg.ctx); err != nil {
			log.Fatalf("POSTGRES_ERROR: %s\n", err.Error())
		}
		log.Println("Postgres connection ready!")
		PgxPool = conn
	}
}

func RedisConnection() Option {
	return func(cfg *Config) {
		RedisCache = initializeRedisClient(cfg.ctx, "CACHE", cfg.RedisDsnURL)
		RedisPublisher = initializeRedisClient(cfg.ctx, "PUBLISHER", cfg.RedisDsnURL)
		RedisSubscriber = initializeRedisClient(cfg.ctx, "SUBSCRIBER", cfg.RedisDsnURL)
		log.Println("Redis connection ready!")
	}
}

func initializeRedisClient(
	ctx context.Context,
	connType, redisDsnURL string,
) *redis.Client {
	opts, err := redis.ParseURL(redisDsnURL)
	if err != nil {
		log.Fatalf("REDIS_%s_ERROR: %s\n", connType, err.Error())
	}
	client := redis.NewClient(opts)
	if err := client.Ping(ctx).Err(); err != nil {
		log.Fatalf("REDIS_%s_ERROR: %s\n", connType, err.Error())
	}
	return client
}
