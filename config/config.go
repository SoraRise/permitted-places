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
		Db   `yaml: "database"`
	}
	App struct {
		Version string `env-required:"true" yaml:"version" env: "APP_NAME"`
	}
	HTTP struct {
		Port string `env-required:"true" yaml:"port" env: "PORT_VALUE"`
	}
	Log struct {
		Level string `env-required:"true" yaml:"log_level" env: "LOG_LEVEL"`
	}
	Db struct {
		Host     string `env-required:"true" yaml:"dbhost" env: "DBHOST"`
		DBPort   string `env-required:"true" yaml:"dbport" env: "PORT"`
		User     string `env-required:"true" yaml:"dbuser" env: "USER"`
		Password string `env-required:"true" yaml:"dbpassword" env:"PASSWORD"`
		DBName   string `env-required:"true" yaml:"dbname" env: "DBNAME"`
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}
	err := cleanenv.ReadConfig("./config/config.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config read error %w", err)
	}
	err = cleanenv.ReadEnv(cfg)

	if err != nil {
		return nil, err
	}
	return cfg, nil
}
