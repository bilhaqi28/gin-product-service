//go:build wireinject
// +build wireinject

package dependecies

import (
	"github.com/bilhaqi28/gin-product-service/cache"
	"github.com/bilhaqi28/gin-product-service/config"
	"github.com/bilhaqi28/gin-product-service/controller"
	"github.com/bilhaqi28/gin-product-service/repository"
	"github.com/bilhaqi28/gin-product-service/service"
	"github.com/google/wire"
)

func NewControllerProductWire() controller.ControllerProduct {
	wire.Build(
		cache.NewProductCache,
		config.NewDB,
		repository.NewRepositoryProduct,
		service.NewServiceProduct,
		controller.NewControllerProduct,
	)
	return nil
}
