package models

type User struct {
	ID       string `json: "id" binding: "required"`
	Email    string `json: "email" binding: "required"`
	Password string `json: "password" binding: "required"`
}
