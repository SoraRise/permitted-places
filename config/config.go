package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		HTTP `yaml:"http"`
		App  `yaml: "app"`
		Log  `yaml:"logger"`
	}
	App struct {
		Name string `env-required:"true" yaml:"name" env: "APP_NAME"`
	}
	// HTTP struct {
	// 	Port string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
	// }

	HTTP struct {
		Port string `env-required:"true" yaml:"port" env: "PORT_VALUE"`
	}
	Log struct {
		Level string `env-required:"true" yaml:"log_level" env: "LOG_LEVEL"`
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}
	err := cleanenv.ReadConfig("./config.yml", cfg)

	if err != nil {
		return nil, fmt.Errorf("config read error %w", err)
	}

	err = cleanenv.ReadEnv(cfg)

	if err != nil {
		return nil, err
	}
	return cfg, nil
}
