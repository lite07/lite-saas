package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lite07/lite-saas/models"
)

// GET /users
// Get all users
func FindUsers(c *gin.Context) {
	var users []models.User
	models.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})
}
