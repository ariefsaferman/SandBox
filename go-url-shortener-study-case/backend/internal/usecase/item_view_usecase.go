package usecase

import (
	"context"
	"url-shortener/internal/model"
	"url-shortener/internal/repository"
	cerrors "url-shortener/internal/utils/errors"
)

type ItemViewUsecase interface {
	CreateOne(ctx context.Context, m model.ItemView) (*model.ItemView, error)
}

type itemViewUsecase struct {
	itemViewRepository repository.ItemViewDBRepository
}

func NewItemViewUsecase(itemViewRepository repository.ItemViewDBRepository) ItemViewUsecase {
	return &itemViewUsecase{
		itemViewRepository: itemViewRepository,
	}
}

func (u *itemViewUsecase) CreateOne(ctx context.Context, m model.ItemView) (*model.ItemView, error) {
	res, err := u.itemViewRepository.CreateOne(ctx, m)
	if err != nil {
		return nil, cerrors.NewErrorWrapper(cerrors.CodeServerError, "failed to create item view from repo", err)
	}

	return res, nil
}
