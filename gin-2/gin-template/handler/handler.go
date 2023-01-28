package handler

import "git.garena.com/sea-labs-id/batch-05/gin-template/usecase"

type Handler struct {
	bookUsecase      usecase.BookUsecase
	borrowingUsecase usecase.BorrowingUsecase
	userUsecase      usecase.UserUsecase
}

type Config struct {
	BookUsecase      usecase.BookUsecase
	BorrowingUsecase usecase.BorrowingUsecase
	UserUsecase      usecase.UserUsecase
}

func New(cfg *Config) *Handler {
	return &Handler{
		bookUsecase:      cfg.BookUsecase,
		borrowingUsecase: cfg.BorrowingUsecase,
		userUsecase:      cfg.UserUsecase,
	}
}
