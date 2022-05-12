package service

import (
	"context"

	"github.com/bilhaqi28/gin-product-service/model/web/request"
	"github.com/bilhaqi28/gin-product-service/model/web/response"
)

type ServiceProduct interface {
	FindAll(ctx context.Context) []response.ProductWeb
	Store(ctx context.Context, request request.CreateProduct) response.ProductWeb
}
