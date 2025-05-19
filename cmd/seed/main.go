package main

import (
	"github.com/lite07/lite-saas/internal/database"
)

func main() {
	database.ConnectDatabase()
	database.SeedDatabase()
}
