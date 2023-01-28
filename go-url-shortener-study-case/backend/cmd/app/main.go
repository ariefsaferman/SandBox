package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"url-shortener/cmd/app/handler"
	"url-shortener/cmd/app/router"
	"url-shortener/internal/component/cache"
	"url-shortener/internal/component/config"
	"url-shortener/internal/component/db"
	"url-shortener/internal/component/taskqueue"
	"url-shortener/internal/constants"
	asynq_repository "url-shortener/internal/repository/asynq"
	pg_repository "url-shortener/internal/repository/postgres"
	"url-shortener/internal/usecase"
	"url-shortener/internal/utils/randomizer"

	"github.com/gin-gonic/gin"
	"github.com/hibiken/asynq"
	"github.com/hibiken/asynqmon"
)

func main() {
	cfg := config.InitConfig()
	db := db.NewPostgresSQLDB(cfg)
	redisCache := cache.NewRedisCache(cfg)
	randomizer := randomizer.NewRandomizer(time.Now().UnixMilli())
	taskqueue := taskqueue.NewAsynqTaskQueueClient(cfg)

	itemDBRepository := pg_repository.NewItemPostgresRepository(db)
	itemQueueRepository := asynq_repository.NewItemAsynqRepository(taskqueue)
	itemViewDBRepository := pg_repository.NewItemViewPostgresRepository(db)

	itemViewUsecase := usecase.NewItemViewUsecase(itemViewDBRepository)
	itemUsecase := usecase.NewItemUsecase(cfg, itemDBRepository, itemQueueRepository, itemViewUsecase, redisCache, randomizer)

	itemHandler := handler.NewItemHandler(itemUsecase)

	if cfg.Application.Environment == constants.EnvironmentProduction {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	r.Use(gin.Recovery())

	// register asyncmon
	h := asynqmon.New(asynqmon.Options{
		RootPath:     "/monitor", // RootPath specifies the root for asynqmon app
		RedisConnOpt: asynq.RedisClientOpt{Addr: fmt.Sprintf("%s:%d", cfg.Redis.Host, cfg.Redis.Port)},
	})
	r.Any("/monitor/*a", gin.WrapH(h))

	// register routes
	router.RegisterItemRoutes(r, itemHandler)

	srvPort := fmt.Sprintf(":%d", cfg.Application.Port)
	srv := &http.Server{
		Addr:    srvPort,
		Handler: r,
	}

	log.Println("Running HTTP server at: ", srvPort)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
}
