package repository

import (
	"context"
	"database/sql"

	"github.com/bilhaqi28/gin-product-service/helper"
	"github.com/bilhaqi28/gin-product-service/model/domain"
)

type RepositoryProductImpl struct {
}

// findAll implements RepositoryProduct
func (*RepositoryProductImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Product {
	SQL := "SELECT * FROM products"
	rows, err := tx.QueryContext(ctx, SQL)
	defer rows.Close()
	helper.PanicIfError(err)
	var products []domain.Product
	for rows.Next() {
		product := domain.Product{}
		err := rows.Scan(&product.Id, &product.BarangKode, &product.BarangNama, &product.BarangDesc, &product.BarangFoto, &product.BarangThumbnail)
		helper.PanicIfError(err)
		products = append(products, product)
	}
	return products
}

func (*RepositoryProductImpl) Store(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product {
	SQL := "INSERT INTO products(barang_kode,barang_nama,barang_desc,barang_thumbnail,barang_foto)VALUES(?,?,?,?,?)"
	vals := []interface{}{}
	vals = append(vals, product.BarangKode, product.BarangNama, product.BarangDesc, product.BarangThumbnail, product.BarangFoto)
	result, err := tx.ExecContext(ctx, SQL, vals...)
	helper.PanicIfError(err)
	id, err := result.LastInsertId()
	helper.PanicIfError(err)
	product.Id = int(id)
	return product
}

func NewRepositoryProduct() RepositoryProduct {
	return &RepositoryProductImpl{}
}
