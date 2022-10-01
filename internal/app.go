package app

import (
	"github.com/SoraRise/permitted-places/config"
	"github.com/SoraRise/permitted-places/pkg/logger"
)

func Run(config *config.Config) {
	log := logger.New(config.Log.Level)

	router.NewRouter()
}
