package main

import (
	"fmt"
	"log"

	"github.com/mohar9h/go-graphql-ws-api/internal/config"
	"github.com/mohar9h/go-graphql-ws-api/internal/database"
	"github.com/mohar9h/go-graphql-ws-api/internal/server"
	"github.com/mohar9h/go-graphql-ws-api/pkg/logging"
)

func main() {
	cfg := config.GetConfig()
	logger := logging.NewZapLogger(cfg)

	err := database.InitRedis(cfg)
	if err != nil {
		logger.Fatal(logging.Redis, logging.Startup, err.Error(), nil)
	}
	defer database.CloseRedis()

	err = database.InitPostgres(cfg)
	if err != nil {
		logger.Fatal(logging.Postgres, logging.Startup, err.Error(), nil)
	}
	defer database.CloseDB()

	r := server.SetupServer(cfg)
	address := fmt.Sprintf(":%s", cfg.Server.Port)
	if err := r.Run(address); err != nil {
		log.Fatalf("failed to run server: %v", err)
		logger.Fatal(logging.Internal, logging.Startup, err.Error(), nil)
	}
}
