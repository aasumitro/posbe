package config

import (
	"database/sql"
	"fmt"
	"github.com/go-redis/redis/v9"
	"github.com/spf13/viper"
	"log"
	"sync"
)

var (
	cfgOnce        sync.Once
	dbOnce         sync.Once
	redisCacheOnce sync.Once

	Cfg        *Config
	Db         *sql.DB
	RedisCache *redis.Client
)

type Config struct {
	AppName           string `mapstructure:"APP_NAME"`
	AppDescription    string `mapstructure:"APP_DESC"`
	AppDebug          bool   `mapstructure:"APP_DEBUG"`
	AppVersion        string `mapstructure:"APP_VERSION"`
	AppUrl            string `mapstructure:"APP_URL"`
	DBDriver          string `mapstructure:"DB_DRIVER"`
	DBDsnUrl          string `mapstructure:"DB_DSN_URL"`
	JWTSecretKey      string `mapstructure:"JWT_SECRET_KEY"`
	JWTLifetime       int    `mapstructure:"JWT_LIFETIME"`
	CacheDriver       string `mapstructure:"CACHE_DRIVER"`
	CacheDsnUrl       string `mapstructure:"CACHE_DSN_URL"`
	CrashReportDriver string `mapstructure:"CRASH_REPORT_DRIVER"`
	CrashReportDsnUrl string `mapstructure:"CRASH_REPORT_DSN_URL"`
}

func LoadConfig() {
	// notify that app try to load config file
	log.Println("Load configuration file . . . .")

	cfgOnce.Do(func() {
		// set config file
		viper.SetConfigFile(".env")
		// find environment file
		viper.AutomaticEnv()
		// error handling for specific case
		if err := viper.ReadInConfig(); err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); ok {
				// Config file not found; ignore error if desired
				panic(".env file not found!, please copy .env.example and paste as .env")
			}

			panic(fmt.Sprintf("ENV_ERROR: %s", err.Error()))
		}
		// notify that config file is ready
		log.Println("configuration file: ready")
		// extract config to struct
		if err := viper.Unmarshal(&Cfg); err != nil {
			panic(fmt.Sprintf("ENV_ERROR: %s", err.Error()))
		}
	})
}
