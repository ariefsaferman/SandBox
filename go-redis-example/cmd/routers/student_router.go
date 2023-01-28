package routers

import (
	"git.garena.com/aldino.rahman/go-redis-example/internal/handlers"
	"git.garena.com/aldino.rahman/go-redis-example/internal/repositories/cache"
	"git.garena.com/aldino.rahman/go-redis-example/internal/usecases"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v9"
)

func InitStudentRouter(r *gin.Engine,rdc *redis.Client) {	
	// init repo
	cacheStdRepository := cache.NewStudentCacheRepository(rdc)

	studentUseCase:= usecases.NewStudentUseCase(&usecases.OptsStudentUseCase{
		CacheRepository: cacheStdRepository,
		// you can add another repo db in here
	})

	h:= handlers.New(&handlers.Opts{
		StudentUseCase: studentUseCase,
	})
	r.GET("/students/:id",h.GetStudentbyId)
	r.POST("/students",h.AddStudent)
}