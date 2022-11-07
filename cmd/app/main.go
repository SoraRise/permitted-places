package main

import (
	"fmt"
	"log"

	"github.com/SoraRise/permitted-places/config"
	app "github.com/SoraRise/permitted-places/internal/app"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}
	defer fmt.Println(cfg.HTTP, cfg.Level, cfg.Portdb, cfg.User, cfg.Password, cfg.Adress)
}
