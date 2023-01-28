package usecase

import (
	"git.garena.com/sea-labs-id/batch-05/gin-template/entity"
	"git.garena.com/sea-labs-id/batch-05/gin-template/repository"
)

type BookUsecase interface {
	GetBooks() ([]*entity.Book, error)
	AddBook(book *entity.Book) (*entity.Book, error)
}

type bookUsecaseImpl struct {
	bookRepository repository.BookRepository
}

type BookConfig struct {
	BookRepository repository.BookRepository
}

func NewBookusecase(cfg *BookConfig) BookUsecase {
	return &bookUsecaseImpl{
		bookRepository: cfg.BookRepository,
	}
}

func (u *bookUsecaseImpl) GetBooks() ([]*entity.Book, error) {
	return u.bookRepository.GetAllBook()
}

func (u *bookUsecaseImpl) AddBook(book *entity.Book) (*entity.Book, error) {
	return u.bookRepository.AddBook(book)
}
