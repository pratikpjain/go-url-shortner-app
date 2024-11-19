package config

import (
	"github.com/spf13/viper"
)

type AppConfig struct {
	AppPort string
}

func LoadConfig() *AppConfig {

	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	viper.AutomaticEnv()

	return &AppConfig{
		AppPort: viper.GetString("APP_PORT"),
	}
}
