package controller

import (
	"net/http"

	"github.com/bilhaqi28/gin-product-service/helper"
	"github.com/bilhaqi28/gin-product-service/model/web/request"
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

func (controller *ControllerProductImpl) Store(c *gin.Context) {
	var request request.CreateProduct
	err := c.ShouldBind(&request)
	if err != nil {
		errorResponse := response.ErrorResponse{
			Code:   http.StatusBadRequest,
			Status: false,
			Error:  err.Error(),
		}
		c.JSON(http.StatusBadRequest, errorResponse)
		return
	}
	// parse file
	filename := helper.UploadFile(c.Request, "thumbnail", "produk")
	request.BarangThumbnail = filename

	productWeb := controller.serviceProduct.Store(c.Request.Context(), request)
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
