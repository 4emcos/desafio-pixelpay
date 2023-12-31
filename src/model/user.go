package model

type User struct {
	Name     string `json:"name" binding:"required,gt=3"`
	Document string `json:"document" binding:"required,len=11,lt=14,numeric"`
	Email    string `json:"email" binding:"required,email"`
}
