package usecase

import (
	"errors"
	"time"

	"git.garena.com/sea-labs-id/batch-05/gin-template/entity"
	"git.garena.com/sea-labs-id/batch-05/gin-template/repository"
)

type BorrowingUsecase interface {
	AddBorrowBook(book *entity.BorrowingRecord) (*entity.BorrowingRecord, error)
	ReturnBorrowBook(id int) (*entity.BorrowingRecord, error)
	GetRecordById(id int) (*entity.BorrowingRecord, error)
}

type borrowingUsecaseImpl struct {
	borrowingRepository repository.BorrowingRepository
}

type BorrowingConfig struct {
	BorrowingRepository repository.BorrowingRepository
}

func NewBorrowingUsecase(cfg *BorrowingConfig) BorrowingUsecase {
	return &borrowingUsecaseImpl{
		borrowingRepository: cfg.BorrowingRepository,
	}
}

func (br *borrowingUsecaseImpl) AddBorrowBook(brecord *entity.BorrowingRecord) (*entity.BorrowingRecord, error) {

	return br.borrowingRepository.AddBorrowBook(brecord)
}

func (br *borrowingUsecaseImpl) ReturnBorrowBook(id int) (*entity.BorrowingRecord, error) {
	record, _ := br.borrowingRepository.GetRecordById(id)
	if record == nil {
		return nil, errors.New("there is no book with that ID")
	}

	if record.Status == "Returned" {
		return nil, errors.New("book already returned")
	}

	record.Status = "Returned"
	record.ReturningDate = time.Now()

	return br.borrowingRepository.ReturnBorrowBook(record)
}

func (br *borrowingUsecaseImpl) GetRecordById(id int) (*entity.BorrowingRecord, error) {
	return br.borrowingRepository.GetRecordById(id)
}
