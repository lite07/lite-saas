package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NotFoundResponse(c *gin.Context, message string) {
	c.JSON(http.StatusNotFound, gin.H{"error": message})
}
