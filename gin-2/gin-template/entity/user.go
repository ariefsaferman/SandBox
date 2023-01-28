package entity

import "gorm.io/gorm"

type User struct {
	Id       int    `json:"id" gorm:"primaryKey"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Password string `json:"password"`
	gorm.Model
}

type UserLoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
