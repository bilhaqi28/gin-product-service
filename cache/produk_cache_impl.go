package cache

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/bilhaqi28/gin-product-service/config"
	"github.com/bilhaqi28/gin-product-service/helper"
	"github.com/bilhaqi28/gin-product-service/model/web/response"
	"github.com/go-redis/redis/v8"
)

type ProductCacheImpl struct {
	redisDB     *redis.Client
	redisConfig config.RedisCache
}

func (cache *ProductCacheImpl) Set(ctx context.Context, key string, product response.ProductWeb, expired time.Duration) error {
	data, err := json.Marshal(product)
	helper.PanicIfError(err)
	err = cache.redisDB.Set(ctx, key, data, 20*time.Second).Err()
	if err != nil {
		println("error set")
		return errors.New(err.Error())
	}
	return nil
}

func (cache *ProductCacheImpl) Get(ctx context.Context, key string) (response.ProductWeb, error) {
	var product response.ProductWeb
	val, err := cache.redisDB.Get(ctx, key).Result()
	if err != nil {
		return product, errors.New("data kosong")
	}
	err = json.Unmarshal([]byte(val), &product)
	helper.PanicIfError(err)
	return product, nil
}

func (cache *ProductCacheImpl) SetAll(ctx context.Context, key string, product []response.ProductWeb, expired time.Duration) error {
	data, err := json.Marshal(product)
	helper.PanicIfError(err)
	err = cache.redisDB.Set(ctx, key, data, 20*time.Second).Err()
	if err != nil {
		println("error set")
		return errors.New(err.Error())
	}
	return nil
}

func (cache *ProductCacheImpl) GetAll(ctx context.Context, key string) ([]response.ProductWeb, error) {
	var product []response.ProductWeb
	val, err := cache.redisDB.Get(ctx, key).Result()
	if err != nil {
		return product, errors.New("data kosong")
	}
	err = json.Unmarshal([]byte(val), &product)
	helper.PanicIfError(err)
	return product, nil
}

func NewProductCache() ProductCache {
	redisObj := config.NewRedis()
	redisClient := redisObj.ConnectRedis()
	return &ProductCacheImpl{
		redisDB:     redisClient,
		redisConfig: redisObj,
	}
}
