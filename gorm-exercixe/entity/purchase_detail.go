package entity

import (
	"time"

	"gorm.io/gorm"
)

type PurchaseDetail struct {
	PurchaseDetailId int `gorm:"primaryKey;column:purchase_detail_id"`
	PurchaseId       int `gorm:"references:PurchaseId"`
	ItemId           int
	Price            float32
	Quantity         int
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt
	// PurchaseHeader PurchaseHeader `gorm:"foreignKey:PurchaseId;references:PurchaseId"`
}
