package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lite07/lite-saas/internal/database"
	"github.com/lite07/lite-saas/internal/routers"
)

func main() {
	r := gin.Default()

	database.ConnectDatabase()

	routers.RegisterUsersRoute(r)
	routers.RegisterSessionsRoute(r)
	routers.RegisterRolesRoute(r)

	r.Run()
}
