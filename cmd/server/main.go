package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/lite07/lite-saas/internal/database"
	"github.com/lite07/lite-saas/internal/routers"
)

func main() {
	r := gin.Default()

	database.ConnectDatabase()

	if os.Getenv("SEED_DATABASE") == "true" {
		database.SeedDatabase()
	}

	routers.RegisterUsersRoute(r)
	routers.RegisterSessionsRoute(r)
	routers.RegisterRolesRoute(r)

	r.Run()
}
