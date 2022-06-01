package service

import (
	"context"
	"database/sql"
	"errors"

	"github.com/bilhaqi28/gin-product-service/helper"
	"github.com/bilhaqi28/gin-product-service/model/domain"
	"github.com/bilhaqi28/gin-product-service/model/web/request"
	"github.com/bilhaqi28/gin-product-service/model/web/response"
	"github.com/bilhaqi28/gin-product-service/repository"
)

type ServiceProductImpl struct {
	repositoryProduct repository.RepositoryProduct
	DB                *sql.DB
}

// Destroy implements ServiceProduct
func (service *ServiceProductImpl) Destroy(ctx context.Context, productId int) string {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	err = service.repositoryProduct.Destroy(ctx, tx, productId)
	if err != nil {
		return err.Error()
	} else {
		return ""
	}
}

// Update implements ServiceProduct
func (service *ServiceProductImpl) Update(ctx context.Context, request request.UpdateProduct, productId int) response.ProductWeb {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	product, err := service.repositoryProduct.FindById(ctx, tx, productId)
	helper.PanicIfError(err)

	// cek apakah ada foto
	if request.BarangThumbnail == "" {
		request.BarangThumbnail = product.BarangThumbnail
	}

	productUpdate := domain.Product{
		BarangKode:      request.BarangKode,
		BarangNama:      request.BarangNama,
		BarangDesc:      request.BarangDesc,
		BarangThumbnail: request.BarangThumbnail,
		BarangFoto:      request.BarangFoto,
		Id:              product.Id,
	}

	result := service.repositoryProduct.Update(ctx, tx, productUpdate)
	return helper.ToProductResponse(result)
}

// FindById implements ServiceProduct
func (service *ServiceProductImpl) FindById(ctx context.Context, productId int) (response.ProductWeb, error) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)
	product, err := service.repositoryProduct.FindById(ctx, tx, productId)
	if err != nil {
		return helper.ToProductResponse(product), errors.New(err.Error())
	}
	return helper.ToProductResponse(product), nil
}

// findAll implements ServiceProduct
func (service *ServiceProductImpl) FindAll(ctx context.Context) []response.ProductWeb {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	products := service.repositoryProduct.FindAll(ctx, tx)
	return helper.ToProductResponses(products)
}

func (service *ServiceProductImpl) Store(ctx context.Context, request request.CreateProduct) response.ProductWeb {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	product := domain.Product{
		BarangKode:      request.BarangKode,
		BarangNama:      request.BarangNama,
		BarangDesc:      request.BarangDesc,
		BarangThumbnail: request.BarangThumbnail,
		BarangFoto:      request.BarangFoto,
	}

	result := service.repositoryProduct.Store(ctx, tx, product)
	return helper.ToProductResponse(result)
}

func NewServiceProduct(repositoryProduct repository.RepositoryProduct, DB *sql.DB) ServiceProduct {
	return &ServiceProductImpl{
		repositoryProduct: repositoryProduct,
		DB:                DB,
	}
}
