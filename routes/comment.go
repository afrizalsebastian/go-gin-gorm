package routes

import (
	comment_controllers "github.com/afrizalsebastian/go-gin-gorm/controllers/comment"
	"github.com/afrizalsebastian/go-gin-gorm/middleware"
	"github.com/gin-gonic/gin"
)

func SetupCommentRoutes(router *gin.RouterGroup) {
	postGroup := router.Group("/post/:postId/comment")
	{
		postGroup.POST("/", middleware.AuthenticationMiddleware, comment_controllers.Create)
		postGroup.GET("/:commentId", middleware.AuthenticationMiddleware, comment_controllers.GetById)
		postGroup.PUT("/:commentId", middleware.AuthenticationMiddleware, comment_controllers.Update)
		postGroup.DELETE("/:commentId", middleware.AuthenticationMiddleware, comment_controllers.Delete)
	}
}
