package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/lite07/lite-saas/internal/requests"
	"github.com/lite07/lite-saas/internal/services"
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
