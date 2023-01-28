package postgres

import (
	"context"
	"errors"
	"url-shortener/internal/model"
	"url-shortener/internal/repository"

	cerrors "url-shortener/internal/utils/errors"

	"github.com/jackc/pgconn"
	"gorm.io/gorm"
)

type itemPostgresRepository struct {
	db *gorm.DB
}

func NewItemPostgresRepository(db *gorm.DB) repository.ItemDBRepository {
	return &itemPostgresRepository{
		db: db,
	}
}

func (r *itemPostgresRepository) GetOneBySlug(ctx context.Context, slug string) (*model.Item, error) {
	var res model.Item

	if err := r.db.WithContext(ctx).Where("slug", slug).First(&res).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, cerrors.NewErrorWrapper(cerrors.CodeServerError, "failed to get one item by slug", err)
	}

	return &res, nil
}

func (r *itemPostgresRepository) GetOneById(ctx context.Context, id int64) (*model.Item, error) {
	var res model.Item

	if err := r.db.WithContext(ctx).First(&res, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, cerrors.NewErrorWrapper(cerrors.CodeServerError, "failed to get one item by id", err)
	}

	return &res, nil
}

func (r *itemPostgresRepository) CreateOne(ctx context.Context, m model.Item) (*model.Item, error) {
	if err := r.db.Create(&m).Error; err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch pgErr.Code {
			case "23505": // duplicate key error
				return nil, cerrors.NewErrorWrapper(cerrors.CodeConflictError, "duplicate key", err)
			}

		}
		return nil, cerrors.NewErrorWrapper(cerrors.CodeServerError, "failed to create an item", err)
	}

	return &m, nil
}

func (r *itemPostgresRepository) IncrementViewCountById(ctx context.Context, id int64) error {
	var i *model.Item

	err := r.db.
		Model(&i).
		Where("id", id).
		Update("view_count", gorm.Expr("view_count + 1")).
		Error

	if err != nil {
		return cerrors.NewErrorWrapper(cerrors.CodeServerError, "failed to create an item", err)
	}

	return nil
}
