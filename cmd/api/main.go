package main

import (
	"fmt"
	"github.com/mohar9h/pecalets_backend/internal/config"
	"github.com/mohar9h/pecalets_backend/internal/server"
	"log"
)

func main() {
	cfg := config.GetConfig()
	r := server.SetupServer()

	address := fmt.Sprintf(":%s", cfg.Server.Port)
	if err := r.Run(address); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
