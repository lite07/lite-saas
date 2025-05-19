package middlewares

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lite07/lite-saas/internal/database"
	"github.com/lite07/lite-saas/internal/models"
	"github.com/lite07/lite-saas/internal/utils"
)

func Authenticate(authorizedRoles []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authToken := c.Request.Header.Values("Authorization")

		if authToken == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Authorization token is missing"})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		var session models.Session
		if err := database.DB.Where("Token = ?", authToken).First(&session).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Authorization token is invalid or expired"})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if session.ExpiredAt.Unix() < time.Now().UTC().Unix() {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Authorization token is invalid or expired"})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		var sessionToken models.SessionToken
		if err := json.Unmarshal([]byte(utils.DecryptString(session.Token)), &sessionToken); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Authorization token is invalid or expired"})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		var user models.User
		if err := database.DB.Find(&user).Where("Email = ?", sessionToken.UserEmail).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Authorization token is invalid or expired"})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		var userRoles []models.UserRole
		if err := database.DB.Joins("Role").Find(&userRoles).Where("UserID = ?", user.ID).Error; err != nil || len(userRoles) == 0 {
			c.JSON(http.StatusForbidden, gin.H{"message": "Authorization token is not authorized for this action."})
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		for _, userRole := range userRoles {
			for _, authorizedRole := range authorizedRoles {
				if authorizedRole == userRole.Role.Name {
					c.Next()
					return
				}
			}
		}

		c.JSON(http.StatusForbidden, gin.H{"message": "Authorization token is not authorized for this action."})
		c.AbortWithStatus(http.StatusForbidden)
	}
}
