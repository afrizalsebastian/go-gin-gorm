package routes

import (
	"github.com/afrizalsebastian/go-gin-gorm/controllers"
	"github.com/gin-gonic/gin"
)

func SetupExampleRoutes(router *gin.RouterGroup) {
	exampleGroup := router.Group("/example")
	{
		exampleGroup.GET("/", controllers.GetExample)
	}
}
