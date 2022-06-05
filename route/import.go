package route

import (
	"github.com/bilhaqi28/gin-product-service/controller"
	"github.com/bilhaqi28/gin-product-service/dependecies"
)

type RouterUse interface {
	ProductController() controller.ControllerProduct
}

type RouterUseImpl struct {
}

func (route *RouterUseImpl) ProductController() controller.ControllerProduct {
	return dependecies.NewControllerProductWire()
}

func NewRouteUse() RouterUse {
	return &RouterUseImpl{}
}
