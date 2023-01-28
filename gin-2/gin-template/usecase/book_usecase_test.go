package usecase_test

import (
	"testing"

	"git.garena.com/sea-labs-id/batch-05/gin-template/entity"
	mocks "git.garena.com/sea-labs-id/batch-05/gin-template/mocks/repository"
	"git.garena.com/sea-labs-id/batch-05/gin-template/usecase"
	"git.garena.com/sea-labs-id/batch-05/gin-template/util"
	"github.com/stretchr/testify/assert"
)

func TestGetBooks(t *testing.T) {
	t.Run("should return list of book when there is no error", func(t *testing.T) {
		var books []*entity.Book
		mockRepo := mocks.NewBookRepository(t)
		uc := usecase.NewBookusecase(&usecase.BookConfig{
			BookRepository: mockRepo,
		})
		mockRepo.On("GetAllBook").Return(books, nil)

		res, err := uc.GetBooks()

		assert.Equal(t, books, res)
		assert.Nil(t, err)
	})

	t.Run("should return error when there is internal server error", func(t *testing.T) {
		mockRepo := mocks.NewBookRepository(t)
		uc := usecase.NewBookusecase(&usecase.BookConfig{
			BookRepository: mockRepo,
		})
		mockRepo.On("GetAllBook").Return(nil, util.ErrInternalServer)

		res, err := uc.GetBooks()

		assert.ErrorIs(t, util.ErrInternalServer, err)
		assert.Nil(t, res)
	})

}

func TestAddBook(t *testing.T) {
	t.Run("should return book when no error occur", func(t *testing.T) {
		value := 1
		book := &entity.Book{
			Title:       "Mesake Bangsaku",
			Description: "kebangsaan",
			Quantity:    &value,
			AuthorId:    1,
		}
		mockRepo := mocks.NewBookRepository(t)
		uc := usecase.NewBookusecase(&usecase.BookConfig{
			BookRepository: mockRepo,
		})
		mockRepo.On("AddBook", book).Return(book, nil)

		res, err := uc.AddBook(book)

		assert.Equal(t, book, res)
		assert.Nil(t, err)
	})

	t.Run("should return error when bad request", func(t *testing.T) {
		value2 := -1
		book := &entity.Book{
			Title:       "Mesake Bangsaku",
			Description: "kebangsaan",
			Quantity:    &value2,
			AuthorId:    1,
		}
		mockRepo := mocks.NewBookRepository(t)
		uc := usecase.NewBookusecase(&usecase.BookConfig{
			BookRepository: mockRepo,
		})
		mockRepo.On("AddBook", book).Return(nil, util.ErrBadRequest)

		res, err := uc.AddBook(book)

		assert.ErrorIs(t, util.ErrBadRequest, err)
		assert.Nil(t, res)
	})
}
