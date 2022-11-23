package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	AppName        string `mapstructure:"APP_NAME"`
	AppDescription string `mapstructure:"APP_DESC"`
	AppDebug       bool   `mapstructure:"APP_DEBUG"`
	AppVersion     string `mapstructure:"APP_VERSION"`
	AppUrl         string `mapstructure:"APP_URL"`
	DBDriver       string `mapstructure:"DB_DRIVER"`
	DBDsnUrl       string `mapstructure:"DB_DSN_URL"`
	JWTSecretKey   string `mapstructure:"JWT_SECRET_KEY"`
	JWTLifetime    string `mapstructure:"JWT_LIFETIME"`
}

func LoadConfig() (cfg *Config, err error) {
	// notify that app try to load config file
	log.Println("Load configuration file . . . .")
	// set config file
	viper.SetConfigFile(".env")
	// find environment file
	viper.AutomaticEnv()
	// error handling for specific case
	if err = viper.ReadInConfig(); err != nil {
		return
	}
	// notify that config file is ready
	log.Println("configuration file: ready")
	// extract config to struct
	err = viper.Unmarshal(&cfg)
	// return value
	return
}
