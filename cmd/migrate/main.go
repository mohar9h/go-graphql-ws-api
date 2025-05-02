package main

import (
	"github.com/mohar9h/go-graphql-ws-api/internal/config"
	"github.com/mohar9h/go-graphql-ws-api/internal/database"
	"github.com/mohar9h/go-graphql-ws-api/internal/database/migrations"
)

func main() {

	cfg := config.GetConfig()

	err := database.InitPostgres(cfg)
	if err != nil {
		panic(err)
	}

	migrations.RunMigrations()
}
