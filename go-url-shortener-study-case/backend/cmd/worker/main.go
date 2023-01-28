package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"
	"url-shortener/cmd/worker/handler"
	"url-shortener/internal/component/cache"
	"url-shortener/internal/component/config"
	"url-shortener/internal/component/db"
	"url-shortener/internal/component/taskqueue"
	"url-shortener/internal/constants"
	asynq_repository "url-shortener/internal/repository/asynq"
	pg_repository "url-shortener/internal/repository/postgres"
	"url-shortener/internal/usecase"
	cerrors "url-shortener/internal/utils/errors"
	clog "url-shortener/internal/utils/log"
	"url-shortener/internal/utils/randomizer"

	"github.com/hibiken/asynq"
)

func main() {
	cfg := config.InitConfig()
	db := db.NewPostgresSQLDB(cfg)
	redisCache := cache.NewRedisCache(cfg)
	taskqueue := taskqueue.NewAsynqTaskQueueClient(cfg)
	randomizer := randomizer.NewRandomizer(time.Now().UnixMilli())

	itemDBRepository := pg_repository.NewItemPostgresRepository(db)
	itemQueueRepository := asynq_repository.NewItemAsynqRepository(taskqueue)
	itemViewDBRepository := pg_repository.NewItemViewPostgresRepository(db)

	itemViewUsecase := usecase.NewItemViewUsecase(itemViewDBRepository)
	itemUsecase := usecase.NewItemUsecase(cfg, itemDBRepository, itemQueueRepository, itemViewUsecase, redisCache, randomizer)

	itemHandler := handler.NewItemHandler(itemUsecase)

	srv := asynq.NewServer(
		asynq.RedisClientOpt{Addr: fmt.Sprintf("%s:%d", cfg.Redis.Host, cfg.Redis.Port)},
		asynq.Config{
			Concurrency: cfg.Asynq.ConcurrentWorker,
			Queues: map[string]int{
				"critical": cfg.Asynq.QueueCritical,
				"default":  cfg.Asynq.QueueDefault,
				"low":      cfg.Asynq.QueueLow,
			},
			ErrorHandler: asynq.ErrorHandlerFunc(func(ctx context.Context, task *asynq.Task, err error) {
				var ew *cerrors.ErrorWrapper
				if errors.As(err, &ew) {
					clog.LogAppErr(ew)
				} else {
					log.Println("error: ", err.Error())
				}

				retried, _ := asynq.GetRetryCount(ctx)
				maxRetry, _ := asynq.GetMaxRetry(ctx)
				if retried >= maxRetry {
					log.Printf("retry exhausted for task: %s", task.Type())
				}
			}),
		},
	)

	mux := asynq.NewServeMux()
	mux.HandleFunc(constants.TaskTypeViewCount, itemHandler.HandleViewCountTask)

	if err := srv.Run(mux); err != nil {
		log.Fatalf("could not run server: %v", err)
	}
}
