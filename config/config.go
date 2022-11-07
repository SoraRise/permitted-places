package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		HTTP           `yaml:"http"`
		App            `yaml: "app"`
		Log            `yaml:"logger"`
		DatabaseConfig `yaml:"dbconfig"`
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
	DatabaseConfig struct {
		Portdb   string `env-required:"true" yaml:"portdb" env: "PORT"`
		User     string `env-required:"true" yaml:"user" env: "USER"`
		Password string `env-required:"true" yaml:"password" env: "PASS"`
		Adress   string `env-required:"true" yaml:"adress" env: "ADDRESS"`
	}
)
type confProfile string

func NewConfig(profile confProfile) (*Config, error) {
	cfg := &Config{}
	var err = cleanenv.ReadConfig("", nil)
	switch profile {
	case "windows":
		err = cleanenv.ReadConfig("./config/config.yml", cfg)
	case "macos":
		err = cleanenv.ReadConfig("./config/configWindows.yml", cfg)
	}
	if err != nil {
		return nil, fmt.Errorf("config read error %w", err)
	}
	err = cleanenv.ReadEnv(cfg)

	if err != nil {
		return nil, err
	}
	return cfg, nil
}
