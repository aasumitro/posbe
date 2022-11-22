package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct{}

func init() {
	loadEnv()
}

func loadEnv() Config {
	log.Println("Load configuration file . . . .")
	// find environment file
	viper.SetConfigFile(`.env`)
	// error handling for specific case
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			panic(".env.example file not found!, please copy .env.example.example and paste as .env.example")
		} else {
			// Config file was found but another error was produced
			panic(err)
		}
	}
	log.Println("configuration file: ready")

	return Config{}
}
