package repository

import (
	"context"
	"database/sql"

	"github.com/bilhaqi28/gin-product-service/model/domain"
)

type RepositoryProduct interface {
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Product
	Store(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product
	FindById(ctx context.Context, tx *sql.Tx, productId int) (domain.Product, error)
	Update(ctx context.Context, tx *sql.Tx, product domain.Product) domain.Product
	Destroy(ctx context.Context, tx *sql.Tx, productId int) error
}
