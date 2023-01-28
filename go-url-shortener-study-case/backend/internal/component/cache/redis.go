package cache

import (
	"fmt"
	"url-shortener/internal/component/config"

	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
)

func newRedisCache(cfg config.Config) *cache.Cache {
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", cfg.Redis.Host, cfg.Redis.Port),
	})

	c := cache.New(&cache.Options{
		Redis: rdb,
	})

	return c
}
