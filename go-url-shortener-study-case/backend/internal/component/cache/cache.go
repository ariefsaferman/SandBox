package cache

import (
	"context"
	"time"
	"url-shortener/internal/component/config"

	"github.com/go-redis/cache/v8"
)

/*
A thin wrapper so that usecase don't have direct dependency to the cache implementation
*/
type Cache interface {
	Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error
	Get(ctx context.Context, key string, value interface{}) error
}

type redisCache struct {
	cache *cache.Cache
}

func NewRedisCache(cfg config.Config) Cache {
	c := newRedisCache(cfg)

	return &redisCache{
		cache: c,
	}
}

func (c *redisCache) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	i := cache.Item{
		Ctx:   ctx,
		Key:   key,
		Value: value,
		TTL:   ttl,
	}

	return c.cache.Set(&i)
}

func (c *redisCache) Get(ctx context.Context, key string, value interface{}) error {
	err := c.cache.Get(ctx, key, value)
	if err == cache.ErrCacheMiss {
		return nil
	}

	return err
}
