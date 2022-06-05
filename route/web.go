package route

import (
	"net/http"

	"github.com/bilhaqi28/gin-product-service/exception"
	"github.com/bilhaqi28/gin-product-service/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	routerUse := NewRouteUse()
	router := gin.Default()
	router.Use(gin.CustomRecovery(exception.InternalServerError))
	router.StaticFS("/public", http.Dir("public"))
	router.Use(middleware.AuthToken)
	CRUDServiceProduct(router, routerUse)
	return router
}

func CRUDServiceProduct(router *gin.Engine, routerUse RouterUse) {
	service := router.Group("/service")
	{
		service.GET("/products", func(c *gin.Context) {
			routerUse.ProductController().FindAll(c)
		})
		service.POST("/product", func(c *gin.Context) {
			routerUse.ProductController().Store(c)
		})
		service.GET("/product/:id", func(c *gin.Context) {
			routerUse.ProductController().FindById(c)
		})
		service.PUT("/product/:id", func(c *gin.Context) {
			routerUse.ProductController().Update(c)
		})
		service.DELETE("/product/:id", func(c *gin.Context) {
			routerUse.ProductController().Destroy(c)
		})
	}
}
