package helper

import (
	"github.com/bilhaqi28/gin-product-service/model/domain"
	"github.com/bilhaqi28/gin-product-service/model/web/response"
)

func ToProductResponse(product domain.Product) response.ProductWeb {
	return response.ProductWeb{
		Id:              product.Id,
		BarangKode:      product.BarangKode,
		BarangNama:      product.BarangNama,
		BarangDesc:      product.BarangDesc,
		BarangThumbnail: product.BarangThumbnail,
		BarangFoto:      product.BarangFoto,
	}
}

func ToProductResponses(products []domain.Product) []response.ProductWeb {
	var productWeb []response.ProductWeb
	for _, product := range products {
		productWeb = append(productWeb, ToProductResponse(product))
	}
	return productWeb
}
