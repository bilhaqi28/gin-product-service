package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/bilhaqi28/gin-product-service/helper"
	"github.com/bilhaqi28/gin-product-service/model/domain"
)

type RepositoryProductImpl struct {
}

// Destroy implements RepositoryProduct
func (*RepositoryProductImpl) Destroy(ctx context.Context, tx *sql.Tx, productId int) error {
	SQL := "DELETE FROM products WHERE id=?"
	rows, err := tx.ExecContext(ctx, SQL, productId)
	helper.PanicIfError(err)
	count, err := rows.RowsAffected()
	helper.PanicIfError(err)
	if count < 1 {
		return errors.New("Product Is Not Found")
	} else {
		return nil
	}
}

// Update implements RepositoryProduct
func (*RepositoryProductImpl) Update(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product {
	SQL := "UPDATE products SET barang_kode=?,barang_nama=?,barang_desc=?,barang_foto=?,barang_thumbnail=? WHERE id=?"
	vals := []interface{}{}
	vals = append(vals, product.BarangKode, product.BarangNama, product.BarangDesc, product.BarangThumbnail, product.BarangFoto, product.Id)
	_, err := tx.ExecContext(ctx, SQL, vals...)
	helper.PanicIfError(err)
	return product

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

func (*RepositoryProductImpl) FindById(ctx context.Context, tx *sql.Tx, productId int) (domain.Product, error) {
	SQL := "SELECT * FROM products WHERE id=?"
	rows, err := tx.QueryContext(ctx, SQL, productId)
	defer rows.Close()
	helper.PanicIfError(err)
	product := domain.Product{}
	if rows.Next() {
		err := rows.Scan(&product.Id, &product.BarangKode, &product.BarangNama, &product.BarangDesc, &product.BarangFoto, &product.BarangThumbnail)
		helper.PanicIfError(err)
		return product, nil
	} else {
		return product, errors.New("Product Is Not Found")
	}

}

func NewRepositoryProduct() RepositoryProduct {
	return &RepositoryProductImpl{}
}
