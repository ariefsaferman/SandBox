package server

import (
	"git.garena.com/sea-labs-id/batch-05/gin-template/handler"
	"git.garena.com/sea-labs-id/batch-05/gin-template/middleware"
	"git.garena.com/sea-labs-id/batch-05/gin-template/usecase"
	"github.com/gin-gonic/gin"
)

type RouterConfig struct {
	BookUsecase      usecase.BookUsecase
	BorrowingUsecase usecase.BorrowingUsecase
	UserUsecase      usecase.UserUsecase
}

func NewRouter(cfg *RouterConfig) *gin.Engine {
	router := gin.Default()
	h := handler.New(&handler.Config{
		BookUsecase:      cfg.BookUsecase,
		BorrowingUsecase: cfg.BorrowingUsecase,
		UserUsecase:      cfg.UserUsecase,
	})

	router.Static("/docs", "swagger-ui")
	router.GET("/books", middleware.AuthTokenMiddleware(), h.GetBooks)
	router.POST("/books", middleware.AuthTokenMiddleware(), h.AddBook)
	router.POST("/borrow", middleware.AuthTokenMiddleware(), h.AddBorrowBook)
	router.PUT("/borrow/:id", middleware.AuthTokenMiddleware(), h.ReturnBorrowBook)
	router.POST("/signin", h.Login)
	return router
}
