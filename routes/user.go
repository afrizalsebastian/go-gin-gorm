package routes

import (
	"github.com/afrizalsebastian/go-gin-gorm/controllers"
	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(router *gin.RouterGroup) {
	userGroup := router.Group("/user")
	{
		userGroup.POST("/register", controllers.CreateUser)
		userGroup.POST("/login", controllers.Login)
	}
}
