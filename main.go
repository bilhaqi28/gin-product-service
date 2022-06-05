package main

import (
	"github.com/bilhaqi28/gin-product-service/route"
)

func main() {
	router := route.NewRouter()
	router.Run(":8080")
}
