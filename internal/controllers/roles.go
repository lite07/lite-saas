package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/lite07/lite-saas/internal/database"
	"github.com/lite07/lite-saas/internal/models"
	"github.com/lite07/lite-saas/internal/requests"
	"github.com/lite07/lite-saas/internal/utils"
)

func GetRoles(c *gin.Context) {
	var roles []models.Role
	database.DB.Find(&roles)

	c.JSON(http.StatusOK, gin.H{"data": roles})
}

func CreateRole(c *gin.Context) {
	// Validate input
	var input requests.CreateRole
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": strings.Split(err.Error(), "\n")})
		return
	}

	role := models.Role{
		Name: input.Name,
	}

	if err := database.DB.Create(&role).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": role})
}

func DeleteRole(c *gin.Context) {
	var role models.Role
	if err := database.DB.Where("ID = ?", c.Param("id")).First(&role).Error; err != nil {
		utils.NotFoundResponse(c, "Role does not exist")
	} else {
		database.DB.Delete(&role)
		c.JSON(http.StatusOK, gin.H{"message": "Successfully deleted"})
	}
}
