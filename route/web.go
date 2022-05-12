package route

import (
	"net/http"

	"github.com/bilhaqi28/gin-product-service/controller"
	"github.com/gin-gonic/gin"
)

func NewRouter(controller controller.MainController) *gin.Engine {
	router := gin.Default()
	router.StaticFS("/public", http.Dir("public"))
	service := router.Group("/service")
	{
		service.GET("/products", func(c *gin.Context) {
			controller.ProductController().FindAll(c)
		})
		service.POST("/product", func(c *gin.Context) {
			controller.ProductController().Store(c)
		})
	}

	return router
}
