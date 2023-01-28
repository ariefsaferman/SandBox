package repository

import (
	"git.garena.com/sea-labs-id/batch-05/gin-template/entity"
	"gorm.io/gorm"
)

type BookRepository interface {
	GetAllBook() ([]*entity.Book, error)
	AddBook(book *entity.Book) (*entity.Book, error)
}

type bookRepositoryImpl struct {
	db *gorm.DB
}

type BookConfig struct {
	DB *gorm.DB
}

func NewBookRepository(cfg *BookConfig) BookRepository {
	return &bookRepositoryImpl{
		db: cfg.DB,
	}
}

func (r *bookRepositoryImpl) GetAllBook() ([]*entity.Book, error) {
	var p []*entity.Book
	res := r.db.Preload("Author").Find(&p)

	if res.Error != nil {
		return nil, res.Error
	}

	return p, nil
}

func (r *bookRepositoryImpl) AddBook(book *entity.Book) (*entity.Book, error) {
	newBook := book
	res := r.db.Create(&newBook)
	if res.Error != nil {
		return nil, res.Error
	}

	return newBook, nil
}
