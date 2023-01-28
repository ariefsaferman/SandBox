package entity

import (
	"time"

	"gorm.io/gorm"
)

type PurchaseHeader struct {
	PurchaseId      int `gorm:"primaryKey;column:purchase_id"`
	InvoiceNo       string
	PurchaseDate    time.Time
	CustomerId      int
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt
	PurchaseDetails []PurchaseDetail `gorm:"foreignKey: PurchaseId"`
}
