package model

import "time"

type Item struct {
	Id        int64     `gorm:"column:id"`
	LongUrl   string    `gorm:"column:long_url"`
	Slug      string    `gorm:"column:slug"`
	Title     string    `gorm:"column:title"`
	Content   string    `gorm:"column:content"`
	ViewCount int64     `gorm:"column:view_count"`
	CreatedAt time.Time `gorm:"column:created_at"`
}
