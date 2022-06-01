package controller

import (
	"database/sql"

	"github.com/bilhaqi28/gin-product-service/dependecies"
)

type MainControllerImpl struct {
	db *sql.DB
}

// productController implements MainController
func (main *MainControllerImpl) ProductController() ControllerProduct {
	service := dependecies.InitializedServer()
	controller := NewControllerProduct(service)
	return controller
}

func NewMainController(db *sql.DB) MainController {
	return &MainControllerImpl{
		db: db,
	}
}
