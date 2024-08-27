package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lite07/lite-saas/src/models"
	"github.com/lite07/lite-saas/src/routers"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	routers.RegisterUsersRoute(r)

	r.Run()
}
