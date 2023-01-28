package server

import (
	"log"

	"git.garena.com/sea-labs-id/batch-05/gin-template/db"
	"git.garena.com/sea-labs-id/batch-05/gin-template/repository"
	"git.garena.com/sea-labs-id/batch-05/gin-template/usecase"
	"github.com/gin-gonic/gin"
)

func createRouter() *gin.Engine {

	bookRepo := repository.NewBookRepository(&repository.BookConfig{
		DB: db.Get(),
	})
	borrowingRepo := repository.NewBorrowingRepository(&repository.BorrowingConfig{
		DB: db.Get(),
	})
	userRepo := repository.NewUserRepository(&repository.UserConfig{
		DB: db.Get(),
	})

	bookUsecase := usecase.NewBookusecase(&usecase.BookConfig{
		BookRepository: bookRepo,
	})
	borrowingUsecase := usecase.NewBorrowingUsecase(&usecase.BorrowingConfig{
		BorrowingRepository: borrowingRepo,
	})
	userUsecase := usecase.NewUserUsecase(&usecase.UserConfig{
		UserRepository: userRepo,
	})

	return NewRouter(&RouterConfig{
		BookUsecase:      bookUsecase,
		BorrowingUsecase: borrowingUsecase,
		UserUsecase:      userUsecase,
	})
}

func Init() {
	r := createRouter()
	err := r.Run()
	if err != nil {
		log.Println("error while running server", err)
		return
	}
}
