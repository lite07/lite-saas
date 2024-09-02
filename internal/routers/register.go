package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/lite07/lite-saas/internal/controllers"
)

func RegisterUsersRoute(r *gin.Engine) {
	r.GET("/api/users", controllers.FindUsers)
	r.GET("/api/users/:id", controllers.GetUser)
	r.DELETE("/api/users/:id", controllers.DeleteUser)
	r.POST("/api/users", controllers.CreateUser)
}

func RegisterSessionsRoute(r *gin.Engine) {
	r.POST("api/sessions", controllers.CreateSession)
	r.POST("api/sessions/invalidate", controllers.InvalidateSession)
}

func RegisterRolesRoute(r *gin.Engine) {
	r.GET("api/roles", controllers.GetRoles)
	r.POST("api/roles", controllers.CreateRole)
	r.DELETE("api/roles/:id", controllers.DeleteRole)
}
