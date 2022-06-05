package controller

import (
	"net/http"
	"strconv"
	"time"

	"github.com/bilhaqi28/gin-product-service/cache"
	"github.com/bilhaqi28/gin-product-service/helper"
	"github.com/bilhaqi28/gin-product-service/model/web/request"
	"github.com/bilhaqi28/gin-product-service/model/web/response"
	"github.com/bilhaqi28/gin-product-service/service"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

type ControllerProductImpl struct {
	serviceProduct service.ServiceProduct
	cacheProduct   cache.ProductCache
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
	// cek terlebihdahulu pada cache jika tidak ada maka baru akses service
	var productWeb response.ProductWeb
	productWeb, err = controller.cacheProduct.Get(c.Request.Context(), "Product"+productId)
	if err == redis.Nil || err != nil {
		productWeb, err = controller.serviceProduct.FindById(c.Request.Context(), id)
		if err == nil {
			err = controller.cacheProduct.Set(c.Request.Context(), "Product"+productId, productWeb, 10)
			helper.PanicIfError(err)
		}
	}
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
	// cek terlebihdahulu pada cache jika tidak ada maka baru akses service
	var productWeb []response.ProductWeb
	productWeb, err := controller.cacheProduct.GetAll(c.Request.Context(), "ProductAll")
	if err == redis.Nil || err != nil {
		productWeb = controller.serviceProduct.FindAll(c.Request.Context())
		err = controller.cacheProduct.SetAll(c.Request.Context(), "ProductAll", productWeb, 20*time.Second)
		helper.PanicIfError(err)
	}
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

func NewControllerProduct(serviceProduct service.ServiceProduct, cache cache.ProductCache) ControllerProduct {
	return &ControllerProductImpl{
		serviceProduct: serviceProduct,
		cacheProduct:   cache,
	}
}
