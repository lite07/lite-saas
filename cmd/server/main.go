package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/lite07/lite-saas/internal/database"
	"github.com/lite07/lite-saas/internal/routers"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()

	database.ConnectDatabase()

	routers.RegisterUsersRoute(r)

	r.Run()
}
