package repository

import (
	"context"
	"database/sql"

	"github.com/vitorlrrcamargo/full-cycle/pos-go-expert/01-go-expert/16-unit-of-work/internal/db"
	"github.com/vitorlrrcamargo/full-cycle/pos-go-expert/01-go-expert/16-unit-of-work/internal/entity"
)

type CategoryRepositoryInterface interface {
	Insert(ctx context.Context, category entity.Category) error
}

type CategoryRepository struct {
	DB      *sql.DB
	Queries *db.Queries
}

func NewCategoryRepository(dtb *sql.DB) *CategoryRepository {
	return &CategoryRepository{
		DB:      dtb,
		Queries: db.New(dtb),
	}
}

func (r *CategoryRepository) Insert(ctx context.Context, category entity.Category) error {
	return r.Queries.CreateCategory(ctx, db.CreateCategoryParams{
		Name: category.Name,
	})
}
