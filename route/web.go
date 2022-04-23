package route

import (
	"github.com/bilhaqi28/gin-product-service/controller"
	"github.com/gin-gonic/gin"
)

func NewRouter(controller controller.MainController) *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/service/product")
	{
		v1.GET("/all", func(c *gin.Context) {
			controller.ProductController().FindAll(c)
		})
	}

	return router
}
