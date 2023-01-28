package entity

import (
	"gorm.io/gorm"
)

type BookRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required,min=10"`
	Quantity    int    `json:"quantity" binding:"required,numeric,gte=0"`
	Cover       string `json:"cover"`
	AuthorId    int    `json:"author_id" binding:"required"`
}

type BookResponse struct {
	Id          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Quantity    int     `json:"quantity"`
	Cover       string  `json:"cover,omitempty"`
	AuthorId    int     `json:"author_id"`
	Author      *Author `json:"author,omitempty"`
	gorm.Model  `json:"-"`
}

type Book struct {
	Id          int     `json:"id" gorm:"primaryKey"`
	Title       string  `json:"title" binding:"required"`
	Description string  `json:"description" binding:"required,min=10"`
	Quantity    *int    `json:"quantity" binding:"required,gte=0"`
	Cover       string  `json:"cover,omitempty"`
	AuthorId    int     `json:"author_id" binding:"required"`
	Author      *Author `json:"author,omitempty" gorm:"reference:AuthorId"`
	gorm.Model  `json:"-"`
}

func (b *Book) ToResponseBook() *BookResponse {
	return &BookResponse{
		Id:          b.Id,
		Title:       b.Title,
		Description: b.Description,
		Quantity:    *b.Quantity,
		Cover:       b.Cover,
		AuthorId:    b.AuthorId,
		Author:      b.Author,
	}
}
