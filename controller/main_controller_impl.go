package controller

import (
	"database/sql"

	"github.com/bilhaqi28/gin-product-service/repository"
	"github.com/bilhaqi28/gin-product-service/service"
)

type MainControllerImpl struct {
	db *sql.DB
}

// productController implements MainController
func (main *MainControllerImpl) ProductController() ControllerProduct {
	repository := repository.NewRepositoryProduct()
	service := service.NewServiceProduct(repository, main.db)
	controller := NewControllerProduct(service)
	return controller
}

func NewMainController(db *sql.DB) MainController {
	return &MainControllerImpl{
		db: db,
	}
}
