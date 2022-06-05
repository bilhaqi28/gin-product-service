package cache

import (
	"context"
	"time"

	"github.com/bilhaqi28/gin-product-service/model/web/response"
)

type ProductCache interface {
	Set(ctx context.Context, key string, product response.ProductWeb, expired time.Duration) error
	Get(ctx context.Context, key string) (response.ProductWeb, error)
	SetAll(ctx context.Context, key string, product []response.ProductWeb, expired time.Duration) error
	GetAll(ctx context.Context, key string) ([]response.ProductWeb, error)
}
