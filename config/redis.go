package config

import (
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisCache struct {
	Host     string
	Password string
	Db       int
	Expires  time.Duration
}

func (cache *RedisCache) ConnectRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cache.Host,
		Password: cache.Password,
		DB:       cache.Db,
	})
	return rdb
}

func NewRedis() RedisCache {
	return RedisCache{
		Host:     "localhost:7001",
		Password: "",
		Db:       0,
		Expires:  10,
	}
}
