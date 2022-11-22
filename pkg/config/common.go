package config

import "github.com/spf13/viper"

func (cfg Config) GetAppName() string {
	return viper.GetString("APP_NAME")
}

func (cfg Config) GetAppDesc() string {
	return viper.GetString("APP_DESC")
}

func (cfg Config) GetAppDebug() bool {
	return viper.GetBool("APP_DEBUG")
}

func (cfg Config) GetAppVersion() string {
	return viper.GetString("APP_VERSION")
}

func (cfg Config) GetAppUrl() string {
	return viper.GetString("APP_URL")
}

func (cfg Config) GetDbDriver() string {
	return viper.GetString("DB_DRIVER")
}

func (cfg Config) GetDbDsnUrl() string {
	return viper.GetString("DB_DSN_URL")
}

func (cfg Config) GetJWTSecretKey() string {
	return viper.GetString("JWT_SECRET_KEY")
}

func (cfg Config) GetJWTLifespan() int {
	return viper.GetInt("JWT_LIFETIME")
}
