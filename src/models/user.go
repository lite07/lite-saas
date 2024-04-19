package models

type User struct {
	ID           uint   `json:"id" gorm:"primary_key"`
	Email        string `json:"email" gorm:"index:idx_email,unique"`
	PasswordHash string `json:"password_hash"`
}
