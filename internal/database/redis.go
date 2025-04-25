package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/mohar9h/go-graphql-ws-api/internal/config"
	"github.com/redis/go-redis/v9"
)

var redisClient *redis.Client

func InitRedis(cfg *config.Config) error {

	redisClient = redis.NewClient(&redis.Options{
		Addr:            fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port),
		Password:        cfg.Redis.Password,
		DialTimeout:     cfg.Redis.DialTimeout * time.Second,
		ReadTimeout:     cfg.Redis.ReadTimeout * time.Second,
		WriteTimeout:    cfg.Redis.WriteTimeout * time.Second,
		PoolSize:        cfg.Redis.PoolSize,
		PoolTimeout:     cfg.Redis.PoolTimeout * time.Second,
		ConnMaxIdleTime: 30 * time.Second,
	})

	_, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		return err
	}

	log.Println("Connected to Redis")

	return nil
}

func GetRedisClient() *redis.Client {
	return redisClient
}

func CloseRedis() error {
	if redisClient != nil {
		return redisClient.Close()
	}
	return nil
}
