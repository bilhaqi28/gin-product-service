package test

import (
	"context"
	"log"
	"testing"

	"github.com/go-redis/redis/v8"
)

// docker run --name redis7 -p 7001:6379 -d redis
func TestRedis(t *testing.T) {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:       "localhost:7001",
		Password:   "",
		DB:         0,
		MaxRetries: 10,
	})
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Print(err.Error())
	} else {
		log.Print("ruuning")
		log.Print(pong)
	}
}
