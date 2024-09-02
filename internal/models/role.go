package models

type Role struct {
	ID   uint   `json:"id" gorm:"primary_key"`
	Name string `json:"name" gorm:"uniqueIndex"`
}
