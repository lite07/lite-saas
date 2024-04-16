package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lite07/lite-saas/controllers"
	"github.com/lite07/lite-saas/models"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	r.GET("/api/users", controllers.FindUsers)

	r.Run()
}
