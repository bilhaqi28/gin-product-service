package repository

import (
	"context"
	"database/sql"

	"github.com/bilhaqi28/gin-product-service/model/domain"
)

type RepositoryProduct interface {
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Product
}
