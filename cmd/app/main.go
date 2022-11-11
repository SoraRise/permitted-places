package main

import (
	"fmt"

	"github.com/SoraRise/permitted-places/internal/config"
	rise "github.com/SoraRise/permitted-places/pkg/client"
)

func main() {
	cfg, err := config.NewConfig("windows")

	if err != nil {
		fmt.Printf("%s\n error top", err)
	}

	fmt.Println(cfg.DatabaseHost, cfg.DatabaseName, cfg.DatabasePort, cfg.DatabaseUserName, cfg.UserPassword, cfg.MaxAttemptsConnections)

	client := rise.NewClient(ctx, *cfg)
}
