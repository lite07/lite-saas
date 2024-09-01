package configs

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

var instance *Manager
var once sync.Once

type Manager struct {
	HashSalt            string
	EncryptionSecretKey string
}

func GetManager() *Manager {
	once.Do(func() {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}

		instance = &Manager{
			HashSalt:            os.Getenv("HASH_SALT"),
			EncryptionSecretKey: os.Getenv("ENCRYPTION_SECRET_KEY"),
		}
	})

	return instance
}
