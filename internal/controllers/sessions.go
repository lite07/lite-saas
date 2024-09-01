package controllers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/lite07/lite-saas/internal/database"
	"github.com/lite07/lite-saas/internal/models"
	"github.com/lite07/lite-saas/internal/requests"
	"github.com/lite07/lite-saas/internal/services"
	"github.com/lite07/lite-saas/internal/utils"
)

func CreateSession(c *gin.Context) {
	// Validate input
	var input requests.CreateUserSession
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": strings.Split(err.Error(), "\n")})
		return
	}

	session, err := services.CreateUserSession(input)
	if err == services.ErrorInvalidCredential {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": session})
}

func InvalidateSession(c *gin.Context) {
	var session models.Session
	sessionToken := c.Request.Header.Get("Authorization")
	fmt.Println(sessionToken)

	if err := database.DB.Where("Token = ?", sessionToken).First(&session).Error; err != nil {
		utils.NotFoundResponse(c, "Session with the specified token is not found")
		return
	}

	if err := database.DB.Delete(&session).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Ok"})
}
