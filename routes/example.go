package routes

import (
	"github.com/afrizalsebastian/go-gin-gorm/controllers"
	"github.com/afrizalsebastian/go-gin-gorm/middleware"
	"github.com/gin-gonic/gin"
)

func SetupExampleRoutes(router *gin.RouterGroup) {
	exampleGroup := router.Group("/example")
	{
		exampleGroup.GET("/", middleware.AuthenticationMiddleware, controllers.GetExample)
		exampleGroup.GET("/error", controllers.GetExampleError)
	}
}
