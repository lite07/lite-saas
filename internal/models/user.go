package models

type User struct {
	ID           uint   `json:"id" gorm:"primary_key"`
	FirstName    string `json:"first_name" gorm:"no null; default:null"`
	LastName     string `json:"last_name"`
	Email        string `json:"email" gorm:"index:idx_email,unique"`
	PasswordHash string `json:"password_hash"`
}
