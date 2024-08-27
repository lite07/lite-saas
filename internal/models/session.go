package models

import "time"

type Session struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	Token     string    `json:"token" gorm:"uniqueIndex"`
	ExpiredAt time.Time `json:"expired_at"`
}
