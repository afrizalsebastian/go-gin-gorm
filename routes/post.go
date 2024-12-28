package routes

import (
	post_controllers "github.com/afrizalsebastian/go-gin-gorm/controllers/post"
	"github.com/afrizalsebastian/go-gin-gorm/middleware"
	"github.com/gin-gonic/gin"
)

func SetupPostRoutes(router *gin.RouterGroup) {
	postGroup := router.Group("/post")
	{
		postGroup.POST("/", middleware.AuthenticationMiddleware, post_controllers.Create)
		postGroup.GET("/:id", middleware.AuthenticationMiddleware, post_controllers.GetById)
	}
}
