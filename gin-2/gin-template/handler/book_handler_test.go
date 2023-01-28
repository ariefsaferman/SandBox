package handler_test

import (
	"bytes"
	"errors"
	"net/http"
	"testing"

	"git.garena.com/sea-labs-id/batch-05/gin-template/entity"
	mocks "git.garena.com/sea-labs-id/batch-05/gin-template/mocks/usecase"
	"git.garena.com/sea-labs-id/batch-05/gin-template/server"
	"git.garena.com/sea-labs-id/batch-05/gin-template/testutils"
	"git.garena.com/sea-labs-id/batch-05/gin-template/util"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/stretchr/testify/assert"
)

func TestHandlerBook(t *testing.T) {
	t.Run("should return list of book with status code 200", func(t *testing.T) {
		value := 2
		books := []*entity.Book{
			{
				Title:       "Buku Memasak",
				Description: "Cara Memasak",
				Cover:       "",
				Quantity:    &value,
				AuthorId:    1,
			},
		}
		expectedBody, _ := json.Marshal(books)
		mockBookUsecase := mocks.NewBookUsecase(t)
		mockBookUsecase.On("GetBooks").Return(books, nil)
		cfg := server.RouterConfig{
			BookUsecase: mockBookUsecase,
		}

		req, _ := http.NewRequest("GET", "/books", nil)
		_, rec := testutils.ServeReq(&cfg, req)

		assert.Equal(t, 200, rec.Code)
		assert.Equal(t, string(expectedBody), rec.Body.String())

	})

	t.Run("should return empty slice when there is no book in database", func(t *testing.T) {
		books := []*entity.Book{}
		expectedBody, _ := json.Marshal(books)
		mockBookUsecase := new(mocks.BookUsecase)
		mockBookUsecase.On("GetBooks").Return(books, nil)
		cfg := server.RouterConfig{
			BookUsecase: mockBookUsecase,
		}
		req, _ := http.NewRequest("GET", "/books", nil)
		_, rec := testutils.ServeReq(&cfg, req)

		assert.Equal(t, 200, rec.Code)
		assert.Equal(t, string(expectedBody), rec.Body.String())
	})

	t.Run("should return error when there is an error in database", func(t *testing.T) {
		uc := mocks.NewBookUsecase(t)
		err := errors.New("")
		uc.On("GetBooks").Return(nil, err)
		expectedRespon, _ := json.Marshal(gin.H{
			"code":    "INTERNAL_SERVER_ERROR",
			"message": err.Error(),
		})

		cfg := server.RouterConfig{
			BookUsecase: uc,
		}

		req, _ := http.NewRequest("GET", "/books", nil)
		_, rec := testutils.ServeReq(&cfg, req)

		assert.Equal(t, string(expectedRespon), rec.Body.String())
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	})
}

func TestAddBook(t *testing.T) {
	value := 5
	t.Run("should return book when create a book", func(t *testing.T) {
		book := &entity.Book{
			Title:       "Hujan",
			Description: "romance-scifi",
			Quantity:    &value,
			AuthorId:    1,
		}
		uc := mocks.NewBookUsecase(t)
		uc.On("AddBook", book).Return(book, nil)
		cfg := server.RouterConfig{
			BookUsecase: uc,
		}

		body, _ := json.Marshal(book)
		req, _ := http.NewRequest("POST", "/books", bytes.NewBuffer(body))
		_, rec := testutils.ServeReq(&cfg, req)

		assert.Equal(t, string(body), rec.Body.String())
		assert.Equal(t, http.StatusOK, rec.Code)
	})

	t.Run("should return error when input same title", func(t *testing.T) {
		book := &entity.Book{
			Title:       "Hujan",
			Description: "romance-scifi",
			Quantity:    &value,
			AuthorId:    1,
		}
		uc := mocks.NewBookUsecase(t)
		uc.On("AddBook", book).Return(nil, errors.New("duplicate key"))
		cfg := server.RouterConfig{
			BookUsecase: uc,
		}
		body, _ := json.Marshal(book)
		expectedError, _ := json.Marshal(gin.H{
			"code":    "SAME_TITLE_ERROR",
			"message": util.ErrBookSameTitle.Error(),
		})

		req, _ := http.NewRequest("POST", "/books", bytes.NewBuffer(body))
		_, rec := testutils.ServeReq(&cfg, req)

		assert.Equal(t, string(expectedError), rec.Body.String())
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})

	t.Run("should return error when bad request", func(t *testing.T) {
		value2 := -1
		book := &entity.Book{
			Title:       "Hujan",
			Description: "romance-scifi",
			Quantity:    &value2,
		}
		uc := mocks.NewBookUsecase(t)
		cfg := server.RouterConfig{
			BookUsecase: uc,
		}
		body, _ := json.Marshal(book)
		expectedError, _ := json.Marshal(gin.H{
			"code":    "BAD_REQUEST",
			"message": util.ErrBadRequest.Error(),
		})

		req, _ := http.NewRequest("POST", "/books", bytes.NewBuffer(body))
		_, rec := testutils.ServeReq(&cfg, req)

		assert.Equal(t, string(expectedError), rec.Body.String())
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	})

	t.Run("should return error when internal server error", func(t *testing.T) {
		book := &entity.Book{
			Title:       "Pohon",
			Description: "romance-scifi",
			Quantity:    &value,
			AuthorId:    1,
		}
		uc := mocks.NewBookUsecase(t)
		uc.On("AddBook", book).Return(nil, errors.New("internal server error"))
		cfg := server.RouterConfig{
			BookUsecase: uc,
		}
		body, _ := json.Marshal(book)
		expectedError, _ := json.Marshal(gin.H{
			"code":    "INTERNAL_SERVER_ERROR",
			"message": util.ErrInternalServer,
		})

		req, _ := http.NewRequest("POST", "/books", bytes.NewBuffer(body))
		_, rec := testutils.ServeReq(&cfg, req)

		assert.Equal(t, string(expectedError), rec.Body.String())
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	})
}
