package entity

import (
	"time"

	"gorm.io/gorm"
)

type ItemCategory struct {
	ItemCategoryId int `gorm:"primaryKey;column:item_category_id"`
	Name           string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt
	Item           []Item `gorm:"foreignKey:ItemCategoryId;references:ItemCategoryId"`
}
