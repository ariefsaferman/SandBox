package entity

import (
	"time"

	"gorm.io/gorm"
)

type Item struct {
	ItemId         int `gorm:"primaryKey;column:item_id"`
	ItemCategoryId int `gorm:"foreignKey;column:item_category_id"`
	Name           string
	Price          float32
	Quantity       int
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt
}
