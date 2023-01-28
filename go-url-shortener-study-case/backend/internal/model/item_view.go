package model

import "time"

type ItemView struct {
	Id        int64     `gorm:"column:id"`
	ItemId    int64     `gorm:"column:item_id"`
	CreatedAt time.Time `gorm:"column:created_at"`
}
