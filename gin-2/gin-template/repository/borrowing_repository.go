package repository

import (
	"errors"

	"git.garena.com/sea-labs-id/batch-05/gin-template/entity"
	"gorm.io/gorm"
)

type BorrowingRepository interface {
	AddBorrowBook(borrow *entity.BorrowingRecord) (*entity.BorrowingRecord, error)
	ReturnBorrowBook(record *entity.BorrowingRecord) (*entity.BorrowingRecord, error)
	GetRecordById(id int) (*entity.BorrowingRecord, error)
}

type borrowRepositoryImpl struct {
	db *gorm.DB
}

type BorrowingConfig struct {
	DB *gorm.DB
}

func NewBorrowingRepository(cfg *BorrowingConfig) BorrowingRepository {
	return &borrowRepositoryImpl{
		db: cfg.DB,
	}
}

func (br *borrowRepositoryImpl) AddBorrowBook(borrow *entity.BorrowingRecord) (*entity.BorrowingRecord, error) {

	//check if the user already borrow the book
	var bu *[]entity.BorrowingRecord
	err := br.db.Model(&bu).Where("user_id = ?", borrow.UserId).Where("book_id = ?", borrow.BookId).Where("status = ?", "Borrowed").First(&bu).Error
	if err == nil {
		return nil, errors.New("user already borrow the book")
	}

	//check if quantity of books == 0
	var book *entity.Book
	err2 := br.db.Model(&book).Where("id = ?", borrow.BookId).Where("quantity >= ?", 1).First(&book).Error
	if err2 != nil {
		return nil, errors.New("book is not available")
	}

	// case if success
	br.db.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(&borrow).Error
		if err != nil {
			return err
		}
		err2 := tx.Model(&entity.Book{}).Where("id = ?", borrow.BookId).Update("quantity", gorm.Expr("quantity - ?", 1)).Error

		if err2 != nil {
			return err2
		}
		return nil
	})
	return borrow, nil
}

func (br *borrowRepositoryImpl) ReturnBorrowBook(record *entity.BorrowingRecord) (*entity.BorrowingRecord, error) {

	err := br.db.Transaction(func(tx *gorm.DB) error {
		err := tx.Save(&record)
		if err.Error != nil {
			return errors.New("interval server error")
		}

		err2 := tx.Model(&entity.Book{}).Where("id = ?", record.BookId).Update("quantity", gorm.Expr("quantity + ?", 1)).Error
		if err2 != nil {
			return errors.New("internal server error")
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	br.db.Preload("User").Preload("Book").First(&record)
	return record, nil
}

func (br *borrowRepositoryImpl) GetRecordById(id int) (*entity.BorrowingRecord, error) {
	var record *entity.BorrowingRecord
	err := br.db.Where("id = ?", id).First(&record).Error
	if err != nil {
		return nil, err
	}
	return record, nil
}
