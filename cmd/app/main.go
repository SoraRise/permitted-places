package main

import (
	"log"

	"github.com/SoraRise/permitted-places/config"
	app "github.com/SoraRise/permitted-places/internal"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	app.Run(cfg)

}
