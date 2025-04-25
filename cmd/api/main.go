package main

import (
	"fmt"
	"log"

	"github.com/mohar9h/go-graphql-ws-api/internal/config"
	"github.com/mohar9h/go-graphql-ws-api/internal/database"
	"github.com/mohar9h/go-graphql-ws-api/internal/server"
)

func main() {
	cfg := config.GetConfig()

	err := database.InitRedis(cfg)
	if err != nil {
		log.Fatalf("failed to init redis: %v", err)
	}
	defer database.CloseRedis()

	err = database.InitPostgres(cfg)
	if err != nil {
		log.Fatalf("failed to init postgres: %v", err)
	}
	defer database.CloseDB()

	r := server.SetupServer()
	address := fmt.Sprintf(":%s", cfg.Server.Port)
	if err := r.Run(address); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
