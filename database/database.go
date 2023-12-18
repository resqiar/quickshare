package database

import (
	"context"
	"log"
	"os"

	redis "github.com/gofiber/storage/redis/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

func InitDatabase() *pgxpool.Pool {
	DSN := os.Getenv("DATABASE_DSN")

	db, err := pgxpool.New(context.Background(), DSN)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(context.Background()); err != nil {
		log.Fatal(err)
	}

	return db
}

var RedisStore *redis.Storage

func InitRedis() {
	REDIS_URL := os.Getenv("REDIS_URL")

	RedisStore = redis.New(redis.Config{
		URL:   REDIS_URL,
		Reset: false,
	})
}
