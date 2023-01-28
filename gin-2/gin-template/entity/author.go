package entity

import "gorm.io/gorm"

type Author struct {
	Id         int    `json:"id" gorm:"primaryKey"`
	Name       string `json:"name"`
	gorm.Model `json:"-"`
}
