package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lite07/lite-saas/models"
	"github.com/lite07/lite-saas/requests"
	"github.com/lite07/lite-saas/utils"
)

// GET /users
// Get all users
func FindUsers(c *gin.Context) {
	var users []models.User
	models.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})
}

// GET /users/:id
// Get specific user
func GetUser(c *gin.Context) {
	var user models.User
	if err := models.DB.Where("ID = ?", c.Param("id")).First(&user).Error; err != nil {
		utils.NotFoundResponse(c, "User does not exist")
	} else {
		c.JSON(http.StatusOK, gin.H{"data": user})
	}
}

// POST /users
// Create new user
func CreateUser(c *gin.Context) {
	// Validate input
	var input requests.CreateUserRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create user
	user := models.User{Email: input.Email, PasswordHash: utils.HashSha256String(input.Password)}
	models.DB.Create(&user)

	c.JSON(http.StatusCreated, gin.H{"data": user})
}

// DEL /users/:id
// Delete specific user
func DeleteUser(c *gin.Context) {
	var user models.User
	if err := models.DB.Where("ID = ?", c.Param("id")).First(&user).Error; err != nil {
		utils.NotFoundResponse(c, "User does not exist")
	} else {
		models.DB.Delete(&user)
		c.JSON(http.StatusOK, gin.H{"message": "Successfully deleted"})
	}
}
