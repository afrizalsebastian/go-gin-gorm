package main

import (
	"log"

	"github.com/afrizalsebastian/go-gin-gorm/config"
	"github.com/afrizalsebastian/go-gin-gorm/middleware"
	"github.com/afrizalsebastian/go-gin-gorm/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file:", err)
		panic(err)
	}

	//Connect To DB
	config.ConnectDatabase()

	r := gin.Default()
	r.Use(middleware.ErrorHandling())

	api := r.Group("/api")
	{
		routes.SetupUserRoutes(api)
		routes.SetupPostRoutes(api)
	}

	r.Run(":8000")
}
