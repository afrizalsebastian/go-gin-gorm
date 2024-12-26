package routes

import (
	user_controllers "github.com/afrizalsebastian/go-gin-gorm/controllers/user"
	"github.com/afrizalsebastian/go-gin-gorm/middleware"
	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(router *gin.RouterGroup) {
	userGroup := router.Group("/user")
	{
		userGroup.GET("/", middleware.AuthenticationMiddleware, user_controllers.Get)
		userGroup.DELETE("/", middleware.AuthenticationMiddleware, user_controllers.Delete)
		userGroup.POST("/register", user_controllers.Create)
		userGroup.POST("/login", user_controllers.Login)
	}
}
