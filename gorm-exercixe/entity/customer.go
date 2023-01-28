package entity

import (
	"time"

	"gorm.io/gorm"
)

type Customers struct {
	CustomerId int `gorm:"primaryKey"`
	Name       string
	Phone      string `gorm:"column: phone_number"`
	Email      string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt
}
