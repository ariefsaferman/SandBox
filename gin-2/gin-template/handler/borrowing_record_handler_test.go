package handler_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	"git.garena.com/sea-labs-id/batch-05/gin-template/entity"
	mocks "git.garena.com/sea-labs-id/batch-05/gin-template/mocks/usecase"
	"git.garena.com/sea-labs-id/batch-05/gin-template/server"
	"git.garena.com/sea-labs-id/batch-05/gin-template/testutils"
	"git.garena.com/sea-labs-id/batch-05/gin-template/util"
	"github.com/stretchr/testify/assert"
)

func TestBorrowRecord(t *testing.T) {
	t.Run("should return borrow record when there is no error", func(t *testing.T) {
		request := entity.BorrowRequest{
			UserId: 1,
			BookId: 5,
		}
		record := request.ToBorrowingRecord()
		expectedResponse, _ := json.Marshal(record)
		mockBorrowUsecase := new(mocks.BorrowingUsecase)
		mockBorrowUsecase.On("AddBorrowBook", record).Return(record, nil)
		cfg := server.RouterConfig{
			BorrowingUsecase: mockBorrowUsecase,
		}
		reqBody := testutils.MakeRequestBody(request)

		req, _ := http.NewRequest("POST", "/borrow", reqBody)
		req.Header.Set("Authorization", "Bearer "+TOKEN)
		_, rec := testutils.ServeReq(&cfg, req)

		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, string(expectedResponse), rec.Body.String())

	})

	t.Run("should return error 400 when bad requesting", func(t *testing.T) {
		request := entity.BorrowRequest{
			UserId: -1,
			BookId: 5,
		}
		error := map[string]string{
			"code":    "BAD_REQUEST",
			"message": "bad request",
		}
		record := request.ToBorrowingRecord()
		expectedResponse, _ := json.Marshal(error)
		mockBorrowUsecase := new(mocks.BorrowingUsecase)
		mockBorrowUsecase.On("AddBorrowBook", record).Return(nil, errors.New("error"))
		cfg := server.RouterConfig{
			BorrowingUsecase: mockBorrowUsecase,
		}
		reqBody := testutils.MakeRequestBody(request)

		req, _ := http.NewRequest("POST", "/borrow", reqBody)
		req.Header.Set("Authorization", "Bearer "+TOKEN)
		_, rec := testutils.ServeReq(&cfg, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, string(expectedResponse), rec.Body.String())

	})

	t.Run("should return error 400 when borrowing out of stock", func(t *testing.T) {
		request := entity.BorrowRequest{
			UserId: 1,
			BookId: 5,
		}
		error := map[string]string{
			"code":    "BAD_REQUEST",
			"message": util.ErrQtyBook.Error(),
		}
		record := request.ToBorrowingRecord()
		expectedResponse, _ := json.Marshal(error)
		mockBorrowUsecase := new(mocks.BorrowingUsecase)
		mockBorrowUsecase.On("AddBorrowBook", record).Return(nil, util.ErrQtyBook)
		cfg := server.RouterConfig{
			BorrowingUsecase: mockBorrowUsecase,
		}
		reqBody := testutils.MakeRequestBody(request)

		req, _ := http.NewRequest("POST", "/borrow", reqBody)
		req.Header.Set("Authorization", "Bearer "+TOKEN)
		_, rec := testutils.ServeReq(&cfg, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, string(expectedResponse), rec.Body.String())

	})

	t.Run("should return error 500 when internal server error", func(t *testing.T) {
		request := entity.BorrowRequest{
			UserId: 1,
			BookId: 5,
		}
		error := map[string]string{
			"code":    "INTERNAL_SERVER_ERROR",
			"message": util.ErrInternalServer.Error(),
		}
		record := request.ToBorrowingRecord()
		expectedResponse, _ := json.Marshal(error)
		mockBorrowUsecase := new(mocks.BorrowingUsecase)
		mockBorrowUsecase.On("AddBorrowBook", record).Return(nil, util.ErrInternalServer)
		cfg := server.RouterConfig{
			BorrowingUsecase: mockBorrowUsecase,
		}
		reqBody := testutils.MakeRequestBody(request)

		req, _ := http.NewRequest("POST", "/borrow", reqBody)
		req.Header.Set("Authorization", "Bearer "+TOKEN)
		_, rec := testutils.ServeReq(&cfg, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Equal(t, string(expectedResponse), rec.Body.String())

	})
}

const TOKEN = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwibmFtZSI6ImFyaWVmIiwiZW1haWwiOiJhcmllZi5zYWZlcm1hbkBzaG9wZWUuY29tIiwiZXhwIjoxNjcwOTM3MTMxLCJpYXQiOjE2NzA5MTkxMzEsImlzcyI6InRlc3QifQ.egg6UW1eEtIjPpc6SKnGpbtpvphHO5ZDgt9fhB826HE"

func TestReturnBook(t *testing.T) {
	t.Run("should return returning record with status code 200 when success returning a book", func(t *testing.T) {
		id := 1
		request := entity.BorrowingRecord{
			ID:     1,
			UserId: 1,
			BookId: 1,
		}
		mockusecase := mocks.NewBorrowingUsecase(t)
		expectedResponse, _ := json.Marshal(request)
		mockusecase.On("ReturnBorrowBook", id).Return(&request, nil)
		mockusecase.On("GetRecordById", id).Return(&request, nil)
		cfg := server.RouterConfig{
			BorrowingUsecase: mockusecase,
		}

		req, _ := http.NewRequest("PUT", "/borrow/1", nil)
		req.Header.Set("Authorization", "Bearer "+TOKEN)
		_, rec := testutils.ServeReq(&cfg, req)

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(expectedResponse), rec.Body.String())

	})

	t.Run("should return error 400 when error bad request", func(t *testing.T) {
		id := 1
		mockusecase := mocks.NewBorrowingUsecase(t)
		mockusecase.On("GetRecordById", id).Return(nil, util.ErrBadRequest)
		expectedResult := map[string]interface{}{
			"code":    "BAD_REQUEST",
			"message": util.ErrBadRequest.Error(),
		}
		cfg := server.RouterConfig{
			BorrowingUsecase: mockusecase,
		}
		expectedResponse, _ := json.Marshal(expectedResult)

		req, _ := http.NewRequest("PUT", "/borrow/1", nil)
		req.Header.Set("Authorization", "Bearer "+TOKEN)
		_, rec := testutils.ServeReq(&cfg, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, string(expectedResponse), rec.Body.String())

	})

	t.Run("should return forbidden when jwt token is not the same", func(t *testing.T) {
		id := 1
		record := &entity.BorrowingRecord{
			ID:     id,
			UserId: 2,
			BookId: 2,
		}
		mockusecase := mocks.NewBorrowingUsecase(t)
		mockusecase.On("GetRecordById", id).Return(record, nil)
		expectedResult := map[string]interface{}{
			"code":    "FORBIDDEN",
			"message": errors.New("forbidden access").Error(),
		}
		cfg := server.RouterConfig{
			BorrowingUsecase: mockusecase,
		}
		expectedResponse, _ := json.Marshal(expectedResult)

		req, _ := http.NewRequest("PUT", "/borrow/1", nil)
		req.Header.Set("Authorization", "Bearer "+TOKEN)
		_, rec := testutils.ServeReq(&cfg, req)

		assert.Equal(t, http.StatusForbidden, rec.Code)
		assert.Equal(t, string(expectedResponse), rec.Body.String())

	})

	t.Run("should return error when error internal server", func(t *testing.T) {
		id := 1
		request := &entity.BorrowingRecord{
			ID:     1,
			UserId: 1,
			BookId: 1,
		}
		response := map[string]interface{}{
			"code":    "INTERNAL_SERVER_ERROR",
			"message": util.ErrInternalServer.Error(),
		}
		mockusecase := mocks.NewBorrowingUsecase(t)
		expectedResponse, _ := json.Marshal(response)
		mockusecase.On("GetRecordById", id).Return(request, nil)
		mockusecase.On("ReturnBorrowBook", id).Return(nil, util.ErrInternalServer)
		cfg := server.RouterConfig{
			BorrowingUsecase: mockusecase,
		}

		req, _ := http.NewRequest("PUT", "/borrow/1", nil)
		req.Header.Set("Authorization", "Bearer "+TOKEN)
		_, rec := testutils.ServeReq(&cfg, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.Equal(t, string(expectedResponse), rec.Body.String())

	})

}
