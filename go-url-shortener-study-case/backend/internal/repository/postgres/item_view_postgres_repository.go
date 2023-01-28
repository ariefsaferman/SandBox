package postgres

import (
	"context"
	"url-shortener/internal/model"
	"url-shortener/internal/repository"
	cerrors "url-shortener/internal/utils/errors"

	"gorm.io/gorm"
)

type itemViewPostgresRepository struct {
	db *gorm.DB
}

func NewItemViewPostgresRepository(db *gorm.DB) repository.ItemViewDBRepository {
	return &itemViewPostgresRepository{
		db: db,
	}
}

func (r *itemViewPostgresRepository) CreateOne(ctx context.Context, m model.ItemView) (*model.ItemView, error) {
	if err := r.db.Create(&m).Error; err != nil {
		return nil, cerrors.NewErrorWrapper(cerrors.CodeServerError, "failed to create an item view", err)
	}

	return &m, nil
}
