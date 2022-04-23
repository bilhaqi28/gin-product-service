package service

import (
	"context"
	"database/sql"

	"github.com/bilhaqi28/gin-product-service/helper"
	"github.com/bilhaqi28/gin-product-service/model/web/response"
	"github.com/bilhaqi28/gin-product-service/repository"
)

type ServiceProductImpl struct {
	repositoryProduct repository.RepositoryProduct
	DB                *sql.DB
}

// findAll implements ServiceProduct
func (service *ServiceProductImpl) FindAll(ctx context.Context) []response.ProductWeb {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	products := service.repositoryProduct.FindAll(ctx, tx)
	return helper.ToProductResponses(products)
}

func NewServiceProduct(repositoryProduct repository.RepositoryProduct, DB *sql.DB) ServiceProduct {
	return &ServiceProductImpl{
		repositoryProduct: repositoryProduct,
		DB:                DB,
	}
}
