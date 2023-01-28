package usecase_test

import (
	"errors"
	"testing"

	"git.garena.com/sea-labs-id/batch-05/gin-template/entity"
	mocks "git.garena.com/sea-labs-id/batch-05/gin-template/mocks/repository"
	"git.garena.com/sea-labs-id/batch-05/gin-template/usecase"
	"github.com/stretchr/testify/assert"
)

func TestUsecaseBorrowBook(t *testing.T) {
	t.Run("should return borrow record when no error occur", func(t *testing.T) {
		borrow := entity.BorrowingRecord{
			UserId: 1,
			BookId: 5,
		}
		mockRepository := mocks.NewBorrowingRepository(t)
		usecase := usecase.NewBorrowingUsecase(&usecase.BorrowingConfig{
			BorrowingRepository: mockRepository,
		})
		mockRepository.On("AddBorrowBook", &borrow).Return(&borrow, nil)

		rec, err := usecase.AddBorrowBook(&borrow)

		assert.NoError(t, err)
		assert.Equal(t, borrow, *rec)
	})

	t.Run("should return error record when borrow book failed", func(t *testing.T) {
		borrow := entity.BorrowingRecord{
			UserId: 1,
			BookId: 5,
		}
		mockRepository := mocks.NewBorrowingRepository(t)
		usecase := usecase.NewBorrowingUsecase(&usecase.BorrowingConfig{
			BorrowingRepository: mockRepository,
		})
		mockRepository.On("AddBorrowBook", &borrow).Return(nil, errors.New("error"))

		_, err := usecase.AddBorrowBook(&borrow)

		assert.Error(t, err)
	})

}
