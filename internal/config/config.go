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

type ConfigWindows struct {
	DatabaseName     string `env-required:"true" yaml:"databaseName" env: "DATABASENAME"`
	DatabasePort     string `env-required:"true" yaml:"portdb" env: "DATABASEPORT"`
	DatabaseUserName string `env-required:"true" yaml:"database_user" env: "USERDATABASE"`
	UserPassword     string `env-required:"true" yaml:"userPassword" env: "USERPASSWORDDB"`
	IPAdress         string `env-required:"true" yaml:"address" env: "IPADRESSDATABASE"`
	adressLocalHost  string `env-required:"false" yaml:"addressLocalHost" env: "TESTADRESSLOCALHOST"`
	DatabasesScheme  string `env-required:"true" yaml:"scheme" env: "DATABASESCHEME"`
}

type confProfile string

func NewConfig(profile confProfile) (*Config, error) {
	cfg := &Config{}
	var err = cleanenv.ReadConfig("", nil)
	switch profile {
	case "windows":
		err = cleanenv.ReadConfig("./config/configMacos.yml", cfgW)
	case "macos":
		err = cleanenv.ReadConfig("./config/configWindows.yml", cfgM)
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
