package repository

import (
	"context"
	"url-shortener/internal/model"
)

type ItemViewDBRepository interface {
	CreateOne(ctx context.Context, m model.ItemView) (*model.ItemView, error)
}
