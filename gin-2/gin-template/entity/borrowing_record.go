package entity

import (
	"time"

	"gorm.io/gorm"
)

type BorrowRequest struct {
	UserId int `json:"user_id" binding:"required,gte=0"`
	BookId int `json:"book_id" binding:"required,gte=0"`
}

type BorrowResponse struct {
	Id            int       `json:"id" gorm:"primaryKey"`
	UserId        int       `json:"user_id"`
	BookId        int       `json:"book_id"`
	Status        string    `json:"status"`
	BorrowingDate time.Time `json:"borrowing_date"`
	ReturningDate time.Time `json:"returning_date"`
	User          *User     `gorm:"foreignKey:UserId" json:"user,omitempty"`
	Book          *Book     `gorm:"foreignKey:BookId" json:"book,omitempty"`
}

type BorrowingRecord struct {
	gorm.Model    `json:"-"`
	ID            int       `json:"id" gorm:"primaryKey"`
	UserId        int       `json:"user_id"`
	BookId        int       `json:"book_id"`
	Status        string    `json:"status"`
	BorrowingDate time.Time `json:"borrowing_date"`
	ReturningDate time.Time `json:"returning_date"`
	User          *User     `gorm:"foreignKey:UserId" json:"user,omitempty"`
	Book          *Book     `gorm:"foreignKey:BookId" json:"book,omitempty"`
}

func (br *BorrowRequest) ToBorrowingRecord() *BorrowingRecord {
	return &BorrowingRecord{
		UserId: br.UserId,
		BookId: br.BookId,
		Status: "Borrowed",
	}
}

func (br *BorrowResponse) FromBorrowingRecord(borrowingRecord *BorrowingRecord) {
	br.Id = borrowingRecord.ID
	br.UserId = borrowingRecord.UserId
	br.BookId = borrowingRecord.BookId
	br.Status = borrowingRecord.Status
}
