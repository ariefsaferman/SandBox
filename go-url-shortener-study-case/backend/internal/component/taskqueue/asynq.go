package taskqueue

import (
	"fmt"
	"url-shortener/internal/component/config"

	"github.com/hibiken/asynq"
)

func NewAsynqTaskQueueClient(cfg config.Config) *asynq.Client {
	client := asynq.NewClient(asynq.RedisClientOpt{
		Addr: fmt.Sprintf("%s:%d", cfg.Redis.Host, cfg.Redis.Port),
	})

	return client
}
