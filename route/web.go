package route

import (
	"net/http"

	"github.com/bilhaqi28/gin-product-service/controller"
	"github.com/bilhaqi28/gin-product-service/exception"
	"github.com/bilhaqi28/gin-product-service/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter(controller controller.MainController) *gin.Engine {
	router := gin.Default()
	router.Use(gin.CustomRecovery(exception.InternalServerError))
	router.StaticFS("/public", http.Dir("public"))
	router.Use(middleware.AuthToken)
	service := router.Group("/service")
	{
		service.GET("/products", func(c *gin.Context) {
			controller.ProductController().FindAll(c)
		})
		service.POST("/product", func(c *gin.Context) {
			controller.ProductController().Store(c)
		})
		service.GET("/product/:id", func(c *gin.Context) {
			controller.ProductController().FindById(c)
		})
		service.PUT("/product/:id", func(c *gin.Context) {
			controller.ProductController().Update(c)
		})
		service.DELETE("/product/:id", func(c *gin.Context) {
			controller.ProductController().Destroy(c)
		})
	}

	return router
}
