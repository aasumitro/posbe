package config

import (
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
		conn := redis.NewClient(&redis.Options{Addr: cfg.RedisDsnURL})
		if err := conn.Ping(cfg.ctx).Err(); err != nil {
			log.Fatalf("REDIS_ERROR: %s\n", err.Error())
		}
		log.Println("Redis connection ready!")
		RdpPool = conn
	}
}
