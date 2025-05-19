package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/lite07/lite-saas/internal/controllers"
	"github.com/lite07/lite-saas/internal/middlewares"
)

var (
	readUser  = []string{"Admin", "SuperAdmin"}
	writeUser = []string{"Admin", "SuperAdmin"}
	readRole  = []string{"Admin", "SuperAdmin"}
	writeRole = []string{"Admin", "SuperAdmin"}
)

func RegisterUsersRoute(r *gin.Engine) {
	r.GET("/api/users", middlewares.Authenticate(readUser), controllers.FindUsers)
	r.GET("/api/users/:id", middlewares.Authenticate(readUser), controllers.GetUser)
	r.DELETE("/api/users/:id", middlewares.Authenticate(writeUser), controllers.DeleteUser)
	r.POST("/api/users", middlewares.Authenticate(writeUser), controllers.CreateUser)
}

func RegisterSessionsRoute(r *gin.Engine) {
	r.POST("api/sessions", controllers.CreateSession)
	r.POST("api/sessions/invalidate", controllers.InvalidateSession)
}

func RegisterRolesRoute(r *gin.Engine) {
	r.GET("api/roles", middlewares.Authenticate(readRole), controllers.GetRoles)
	r.POST("api/roles", middlewares.Authenticate(writeRole), controllers.CreateRole)
	r.DELETE("api/roles/:id", middlewares.Authenticate(writeRole), controllers.DeleteRole)
}
