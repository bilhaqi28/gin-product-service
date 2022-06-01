//go:build wireinject
// +build wireinject

package dependecies

import (
	"github.com/bilhaqi28/gin-product-service/config"
	"github.com/bilhaqi28/gin-product-service/repository"
	"github.com/bilhaqi28/gin-product-service/service"
	"github.com/google/wire"
)

func InitializedServer() service.ServiceProduct {
	wire.Build(
		config.NewDB,
		repository.NewRepositoryProduct,
		service.NewServiceProduct,
	)
	return nil
}
