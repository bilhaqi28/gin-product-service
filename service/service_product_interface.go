package service

import (
	"context"

	"github.com/bilhaqi28/gin-product-service/model/web/response"
)

type ServiceProduct interface {
	FindAll(ctx context.Context) []response.ProductWeb
}
