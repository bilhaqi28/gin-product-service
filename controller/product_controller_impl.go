package controller

import (
	"net/http"

	"github.com/bilhaqi28/gin-product-service/model/web/response"
	"github.com/bilhaqi28/gin-product-service/service"
	"github.com/gin-gonic/gin"
)

type ControllerProductImpl struct {
	serviceProduct service.ServiceProduct
}

// FindAll implements ControllerProduct
func (controller *ControllerProductImpl) FindAll(c *gin.Context) {
	productWeb := controller.serviceProduct.FindAll(c.Request.Context())
	apiResponse := response.ApiResponse{
		Code:   http.StatusOK,
		Status: true,
		Data:   productWeb,
	}
	c.JSON(http.StatusOK, apiResponse)
}

func NewControllerProduct(serviceProduct service.ServiceProduct) ControllerProduct {
	return &ControllerProductImpl{
		serviceProduct: serviceProduct,
	}
}
