package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/lite07/lite-saas/internal/database"
	"github.com/lite07/lite-saas/internal/dto"
	"github.com/lite07/lite-saas/internal/models"
	"github.com/lite07/lite-saas/internal/requests"
	"github.com/lite07/lite-saas/internal/utils"
)

// GET /users
// Get all users
func FindUsers(c *gin.Context) {
	var users []models.User
	database.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": dto.UserDtoFromEntites(users)})
}

// GET /users/:id
// Get specific user
func GetUser(c *gin.Context) {
	var user models.User
	if err := database.DB.Where("ID = ?", c.Param("id")).First(&user).Error; err != nil {
		utils.NotFoundResponse(c, "User does not exist")
	} else {
		c.JSON(http.StatusOK, gin.H{"data": dto.UserDtoFromEntity(user)})
	}
}

// POST /users
// Create new user
func CreateUser(c *gin.Context) {
	// Validate input
	var input requests.CreateUser
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": strings.Split(err.Error(), "\n")})
		return
	}

	// Create user
	user := models.User{
		FirstName:    input.FirstName,
		LastName:     input.LastName,
		Email:        input.Email,
		PasswordHash: utils.HashSha256String(input.Password)}

	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": dto.UserDtoFromEntity(user)})
}

// DEL /users/:id
// Delete specific user
func DeleteUser(c *gin.Context) {
	var user models.User
	if err := database.DB.Where("ID = ?", c.Param("id")).First(&user).Error; err != nil {
		utils.NotFoundResponse(c, "User does not exist")
	} else {
		database.DB.Delete(&user)
		c.JSON(http.StatusOK, gin.H{"message": "Successfully deleted"})
	}
}
