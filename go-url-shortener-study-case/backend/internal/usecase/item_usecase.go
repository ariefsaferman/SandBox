package usecase

import (
	"context"
	"errors"
	"url-shortener/internal/component/cache"
	"url-shortener/internal/component/config"
	"url-shortener/internal/constants"
	"url-shortener/internal/model"
	"url-shortener/internal/repository"
	ccache "url-shortener/internal/utils/cache"
	cerrors "url-shortener/internal/utils/errors"
	clog "url-shortener/internal/utils/log"
	"url-shortener/internal/utils/randomizer"
)

type ItemUsecase interface {
	GetOneBySlug(ctx context.Context, encodedId string) (*model.Item, error)
	GetOneBySlugWithJob(ctx context.Context, slug string) (*model.Item, error)
	GetOneCachedBySlug(ctx context.Context, encodedId string) (*model.Item, error)
	GetOneCachedBySlugWithJob(ctx context.Context, slug string) (*model.Item, error)
	CreateOne(ctx context.Context, m model.Item) (*model.Item, error)
	IncrementViewCountById(ctx context.Context, id int64) error
}

/*
Naive implementation since here it don't use any hashing or other techniques to optimize the slug creation

Just a simple random string then we check if it exists in the database
*/

type itemUsecase struct {
	cfg                 config.Config
	itemDBRepository    repository.ItemDBRepository
	ItemQueueRepository repository.ItemQueueRepository
	itemViewUsecase     ItemViewUsecase
	cache               cache.Cache
	randomizer          randomizer.Randomizer
}

func NewItemUsecase(
	cfg config.Config,
	itemDBRepository repository.ItemDBRepository,
	ItemQueueRepository repository.ItemQueueRepository,
	itemViewUsecase ItemViewUsecase,
	cache cache.Cache,
	randomizer randomizer.Randomizer,
) ItemUsecase {
	return &itemUsecase{
		cfg:                 cfg,
		itemDBRepository:    itemDBRepository,
		ItemQueueRepository: ItemQueueRepository,
		itemViewUsecase:     itemViewUsecase,
		cache:               cache,
		randomizer:          randomizer,
	}
}

func (u *itemUsecase) GetOneBySlug(ctx context.Context, slug string) (*model.Item, error) {
	res, err := u.itemDBRepository.GetOneBySlug(ctx, slug)
	if err != nil {
		return nil, cerrors.NewErrorWrapper(cerrors.CodeServerError, "failed to get one item by slug from repo", err)
	}
	if res == nil {
		return nil, cerrors.NewErrorWrapper(cerrors.CodeNotFoundError, "item not found", nil)
	}

	err = u.IncrementViewCountById(ctx, res.Id)
	if err != nil {
		return nil, cerrors.NewErrorWrapper(cerrors.CodeServerError, "failed to increment view count by id from repo", err)
	}

	_, err = u.itemViewUsecase.CreateOne(ctx, model.ItemView{ItemId: res.Id})
	if err != nil {
		return nil, cerrors.NewErrorWrapper(cerrors.CodeServerError, "failed to create one item view from repo", err)
	}

	return res, err
}

func (u *itemUsecase) GetOneBySlugWithJob(ctx context.Context, slug string) (*model.Item, error) {
	res, err := u.itemDBRepository.GetOneBySlug(ctx, slug)
	if err != nil {
		return nil, cerrors.NewErrorWrapper(cerrors.CodeServerError, "failed to get one item by slug from repo", err)
	}
	if res == nil {
		return nil, cerrors.NewErrorWrapper(cerrors.CodeNotFoundError, "item not found", nil)
	}

	err = u.ItemQueueRepository.IncrementViewCountById(ctx, res.Id)
	if err != nil {
		return nil, cerrors.NewErrorWrapper(cerrors.CodeServerError, "failed to dispatch increment item view count by id", err)
	}

	return res, err
}

func (u *itemUsecase) GetOneCachedBySlug(ctx context.Context, slug string) (*model.Item, error) {
	var res *model.Item

	cacheKey := ccache.BuildCommonKey(constants.CacheKeyItems, slug)

	err := u.cache.Get(ctx, cacheKey, &res)
	if err != nil {
		clog.LogAppErr(cerrors.NewErrorWrapper(cerrors.CodeServerError, "failed to get one item by slug from cache", err))
	}

	if res == nil {
		res, err = u.itemDBRepository.GetOneBySlug(ctx, slug)

		if err != nil {
			return nil, cerrors.NewErrorWrapper(cerrors.CodeServerError, "failed to get one item by slug from repo", err)
		}
		if res == nil {
			return nil, cerrors.NewErrorWrapper(cerrors.CodeNotFoundError, "item not found", nil)
		}

		err = u.cache.Set(ctx, cacheKey, res, constants.CacheOneDay)
		if err != nil {
			clog.LogAppErr(cerrors.NewErrorWrapper(cerrors.CodeNotFoundError, "cannot set item on redis cache", err))
		}
	}

	// increment view count logic, no checking
	err = u.IncrementViewCountById(ctx, res.Id)
	if err != nil {
		return nil, cerrors.NewErrorWrapper(cerrors.CodeServerError, "failed to increment view count by id from repo", err)
	}

	_, err = u.itemViewUsecase.CreateOne(ctx, model.ItemView{ItemId: res.Id})
	if err != nil {
		return nil, cerrors.NewErrorWrapper(cerrors.CodeServerError, "failed to create one item view from repo", err)
	}

	return res, err
}

func (u *itemUsecase) GetOneCachedBySlugWithJob(ctx context.Context, slug string) (*model.Item, error) {
	var res *model.Item

	cacheKey := ccache.BuildCommonKey(constants.CacheKeyItems, slug)

	err := u.cache.Get(ctx, cacheKey, &res)
	if err != nil {
		clog.LogAppErr(cerrors.NewErrorWrapper(cerrors.CodeServerError, "failed to get one item by slug from cache", err))
	}

	if res == nil {
		res, err = u.itemDBRepository.GetOneBySlug(ctx, slug)

		if err != nil {
			return nil, cerrors.NewErrorWrapper(cerrors.CodeServerError, "failed to get one item by slug from repo", err)
		}
		if res == nil {
			return nil, cerrors.NewErrorWrapper(cerrors.CodeNotFoundError, "item not found", nil)
		}

		err = u.cache.Set(ctx, cacheKey, res, constants.CacheOneDay)
		if err != nil {
			clog.LogAppErr(cerrors.NewErrorWrapper(cerrors.CodeNotFoundError, "cannot set item on redis cache", err))
		}
	}

	err = u.ItemQueueRepository.IncrementViewCountById(ctx, res.Id)
	if err != nil {
		return nil, cerrors.NewErrorWrapper(cerrors.CodeServerError, "failed to dispatch increment item view count by id", err)
	}

	return res, err
}

func (u *itemUsecase) CreateOne(ctx context.Context, m model.Item) (*model.Item, error) {

	var (
		res *model.Item
		err error
	)

	if m.Title == "" {
		return nil, cerrors.NewErrorWrapper(cerrors.CodeClientError, "title is required", nil)
	}

	if m.Content == "" {
		return nil, cerrors.NewErrorWrapper(cerrors.CodeClientError, "content is required", nil)
	}

	if m.LongUrl == "" {
		return nil, cerrors.NewErrorWrapper(cerrors.CodeClientError, "long url is required", nil)
	}

	if m.Slug == "" {
		res, err = u.createOneWithRandomSlug(ctx, m)
		if err != nil {
			return nil, err
		}
	} else {
		res, err = u.createOneCustomSlug(ctx, m)
		if err != nil {
			return nil, err
		}
	}

	return res, nil
}

func (u *itemUsecase) createOneWithRandomSlug(ctx context.Context, m model.Item) (*model.Item, error) {
	var (
		res *model.Item
	)

	isCreated := false

	for !isCreated {
		m.Slug = u.randomizer.RandomStr(u.cfg.Application.SlugLength)

		chk, err := u.itemDBRepository.GetOneBySlug(ctx, m.Slug)
		if err != nil {
			return nil, cerrors.NewErrorWrapper(cerrors.CodeServerError, "failed to get one item by slug from repo", err)
		}
		if chk != nil {
			continue
		}

		res, err = u.itemDBRepository.CreateOne(ctx, m)
		// TODO: clean error handling
		if err != nil {
			var appErr *cerrors.ErrorWrapper
			if errors.As(err, &appErr) {
				if appErr.Code == cerrors.CodeConflictError {
					continue
				}
			}
			return nil, cerrors.NewErrorWrapper(cerrors.CodeServerError, "failed to create item from repo", err)
		}

		isCreated = true
	}

	return res, nil
}

func (u *itemUsecase) createOneCustomSlug(ctx context.Context, m model.Item) (*model.Item, error) {
	if m.Slug == "" {
		return nil, cerrors.NewErrorWrapper(cerrors.CodeClientError, "long url is required", nil)
	}

	res, err := u.itemDBRepository.CreateOne(ctx, m)
	if err != nil {
		var appErr *cerrors.ErrorWrapper

		// TODO: clean error handling
		if errors.As(err, &appErr) {
			if appErr.Code == cerrors.CodeConflictError {
				return nil, cerrors.NewErrorWrapper(cerrors.CodeConflictError, "duplicate slug", nil)
			}
		}
		return nil, cerrors.NewErrorWrapper(cerrors.CodeServerError, "failed to create item from repo", err)
	}

	return res, nil
}

func (u *itemUsecase) IncrementViewCountById(ctx context.Context, id int64) error {
	err := u.itemDBRepository.IncrementViewCountById(ctx, id)
	if err != nil {
		return cerrors.NewErrorWrapper(cerrors.CodeServerError, "error increment view count", err)
	}

	_, err = u.itemViewUsecase.CreateOne(ctx, model.ItemView{ItemId: id})
	if err != nil {
		return cerrors.NewErrorWrapper(cerrors.CodeServerError, "failed to create one item view from repo", err)
	}

	return nil
}
