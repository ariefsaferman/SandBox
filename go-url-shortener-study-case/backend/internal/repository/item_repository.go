package repository

import (
	"context"
	"url-shortener/internal/model"
)

type ItemDBRepository interface {
	GetOneBySlug(ctx context.Context, slug string) (*model.Item, error)
	GetOneById(ctx context.Context, id int64) (*model.Item, error)
	CreateOne(ctx context.Context, m model.Item) (*model.Item, error)
	IncrementViewCountById(ctx context.Context, id int64) error
}

type ItemQueueRepository interface {
	IncrementViewCountById(ctx context.Context, id int64) error
}
