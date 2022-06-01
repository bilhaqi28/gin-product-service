package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/bilhaqi28/gin-product-service/helper"
	"github.com/bilhaqi28/gin-product-service/model/web/request"
	"github.com/bilhaqi28/gin-product-service/model/web/response"
	"github.com/bilhaqi28/gin-product-service/service"
	"github.com/gin-gonic/gin"
)

type ControllerProductImpl struct {
	serviceProduct service.ServiceProduct
}

// Destroy implements ControllerProduct
func (controller *ControllerProductImpl) Destroy(c *gin.Context) {
	productId := c.Param("id")
	id, err := strconv.Atoi(productId)
	helper.PanicIfError(err)
	result := controller.serviceProduct.Destroy(c.Request.Context(), id)
	if result == "" {
		apiResponse := response.ApiResponse{
			Code:   http.StatusOK,
			Status: true,
			Data:   "",
		}
		c.JSON(http.StatusOK, apiResponse)
	} else {
		errorResponse := response.ErrorResponse{
			Code:   http.StatusBadRequest,
			Status: false,
			Error:  result,
		}
		c.JSON(http.StatusBadRequest, errorResponse)
	}
}

// Update implements ControllerProduct
func (controller *ControllerProductImpl) Update(c *gin.Context) {
	productId := c.Param("id")
	id, err := strconv.Atoi(productId)
	helper.PanicIfError(err)

	var request request.UpdateProduct
	err = c.ShouldBind(&request)
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
	filename, err := helper.UploadFile(c.Request, "thumbnail", "produk")
	if err == nil {
		request.BarangThumbnail = filename
	}
	productWeb := controller.serviceProduct.Update(c.Request.Context(), request, id)
	apiResponse := response.ApiResponse{
		Code:   http.StatusOK,
		Status: true,
		Data:   productWeb,
	}

	c.JSON(http.StatusOK, apiResponse)

}

// FindById implements ControllerProduct
func (controller *ControllerProductImpl) FindById(c *gin.Context) {
	productId := c.Param("id")
	id, err := strconv.Atoi(productId)
	helper.PanicIfError(err)
	productWeb, err := controller.serviceProduct.FindById(c.Request.Context(), id)
	if err != nil {
		errorResponse := response.ErrorResponse{
			Code:   http.StatusBadRequest,
			Status: false,
			Error:  err.Error(),
		}
		c.JSON(http.StatusBadRequest, errorResponse)
	} else {
		apiResponse := response.ApiResponse{
			Code:   http.StatusOK,
			Status: true,
			Data:   productWeb,
		}
		c.JSON(http.StatusOK, apiResponse)
	}
}

// FindAll implements ControllerProduct
func (controller *ControllerProductImpl) FindAll(c *gin.Context) {
	log.Print(c.Keys["token_jwt"])
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
	filename, _ := helper.UploadFile(c.Request, "thumbnail", "produk")
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
