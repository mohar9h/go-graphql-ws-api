package database

import (
	"fmt"
	"log"
	"time"

	"github.com/mohar9h/go-graphql-ws-api/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbClient *gorm.DB

func InitPostgres(cfg *config.Config) error {
	var err error
	connection := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=Asia/Tehran", cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.Username, cfg.Postgres.Password, cfg.Postgres.Database, cfg.Postgres.SslMode)

	dbClient, err = gorm.Open(postgres.Open(connection), &gorm.Config{})
	if err != nil {
		return err
	}

	sqlDb, _ := dbClient.DB()

	err = sqlDb.Ping()
	if err != nil {
		return err
	}

	sqlDb.SetMaxIdleConns(cfg.Postgres.MaxIdleConns)
	sqlDb.SetMaxOpenConns(cfg.Postgres.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(cfg.Postgres.ConnMaxLifetime * time.Minute)

	log.Println("Connected to Postgres")

	return nil

}

func GetDB() *gorm.DB {
	return dbClient
}

func CloseDB() {
	sqlDb, _ := dbClient.DB()
	sqlDb.Close()
}
