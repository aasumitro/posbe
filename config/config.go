package config

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

var (
	configSingleton, postgresSingleton,
	redisSingleton, engineOnce sync.Once

	Instance     *Config
	PostgresPool *sql.DB
	RedisPool    *redis.Client
	GinEngine    *gin.Engine
)

type Config struct {
	ctx context.Context

	AppName    string `mapstructure:"APP_NAME"`
	AppPort    string `mapstructure:"APP_PORT"`
	AppDebug   bool   `mapstructure:"APP_DEBUG"`
	AppVersion string `mapstructure:"APP_VERSION"`

	PostgresDsnURL string `mapstructure:"POSTGRES_DSN_URL"`
	RedisDsnURL    string `mapstructure:"REDIS_DSN_URL"`
	SentryDsnURL   string `mapstructure:"SENTRY_DSN_URL"`

	JWTSecretKey string `mapstructure:"JWT_SECRET_KEY"`
	JWTLifetime  int    `mapstructure:"JWT_LIFETIME"`
}

type Option func(cfg *Config)

func LoadWith(
	ctx context.Context,
	options ...Option,
) {
	configSingleton.Do(func() {
		// notify that app try to load config file
		log.Println("Load configuration file . . . .")
		// error handling for a specific case
		if err := viper.ReadInConfig(); err != nil {
			var configFileNotFoundError viper.ConfigFileNotFoundError
			if errors.As(err, &configFileNotFoundError) {
				// Config file not found; ignore error if desired
				log.Fatal(".env file not found!, please copy .env.example and paste as .env")
			}
			log.Fatalf("ENV_ERROR: %s", err.Error())
		}
		// notify that the config file is ready
		log.Println("configuration file: ready")
		// extract config to struct
		if err := viper.Unmarshal(&Instance); err != nil {
			log.Fatalf("ENV_ERROR: %s", err.Error())
		}
		// set context
		Instance.ctx = ctx
		// set options
		for _, opt := range options {
			opt(Instance)
		}
	})
}
