package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lite07/lite-saas/internal/models"
	"github.com/lite07/lite-saas/internal/routers"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	routers.RegisterUsersRoute(r)

	r.Run()
}
