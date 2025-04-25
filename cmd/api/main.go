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

	database.InitRedis()
	defer database.CloseRedis()

	r := server.SetupServer()
	address := fmt.Sprintf(":%s", cfg.Server.Port)
	if err := r.Run(address); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
