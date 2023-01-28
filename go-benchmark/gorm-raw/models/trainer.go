package models

type Trainer struct {
	ID   uint   `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}
