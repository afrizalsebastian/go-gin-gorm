package main

import (
	"github.com/afrizalsebastian/go-gin-gorm/config"
	"github.com/afrizalsebastian/go-gin-gorm/middleware"
	"github.com/afrizalsebastian/go-gin-gorm/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()

	r := gin.Default()
	r.Use(middleware.ErrorHandling())

	api := r.Group("/api")
	{
		routes.SetupExampleRoutes(api)
		routes.SetupUserRoutes(api)
	}

	r.Run(":8000")
}
