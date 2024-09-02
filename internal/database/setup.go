package database

import (
	"github.com/lite07/lite-saas/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("./internal/database/development.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	err = database.AutoMigrate(&models.User{})
	if err != nil {
		return
	}
	err = database.AutoMigrate(&models.Session{})
	if err != nil {
		return
	}
	err = database.AutoMigrate(&models.Role{})
	if err != nil {
		return
	}
	err = database.AutoMigrate(&models.UserRole{})
	if err != nil {
		return
	}

	DB = database
}
