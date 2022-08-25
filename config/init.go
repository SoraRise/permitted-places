package config

import (
	"github.com/spf13/viper"
)

func Init() error {
	viper.AddConfigPath("./config.yml")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
