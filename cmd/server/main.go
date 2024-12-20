package main

import (
	"github.com/afrizalsebastian/go-gin-gorm/config"
	"github.com/afrizalsebastian/go-gin-gorm/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDatabase()

	r := gin.Default()
	api := r.Group("/api")
	{
		routes.SetupExampleRoutes(api)
	}

	r.Run(":8000")
}
