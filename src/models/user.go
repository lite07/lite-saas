package models

type User struct {
	ID           uint   `json:"id" gorm:"primary_key"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
}
