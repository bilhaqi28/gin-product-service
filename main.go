package main

import (
	"github.com/bilhaqi28/gin-product-service/config"
	"github.com/bilhaqi28/gin-product-service/controller"
	"github.com/bilhaqi28/gin-product-service/route"
)

func main() {
	db := config.NewDB()
	// product
	controller := controller.NewMainController(db)
	router := route.NewRouter(controller)
	router.Run(":8080")
}
