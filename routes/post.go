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
		postGroup.GET("/", middleware.AuthenticationMiddleware, post_controllers.Get)
		postGroup.GET("/:postId", middleware.AuthenticationMiddleware, post_controllers.GetById)
		postGroup.PUT("/:postId", middleware.AuthenticationMiddleware, post_controllers.Update)
		postGroup.DELETE("/:postId", middleware.AuthenticationMiddleware, post_controllers.Delete)
	}
}
