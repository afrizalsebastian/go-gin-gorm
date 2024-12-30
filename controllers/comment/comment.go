package comment_controllers

import (
	"net/http"
	"strconv"

	"github.com/afrizalsebastian/go-gin-gorm/controllers"
	"github.com/afrizalsebastian/go-gin-gorm/dtos"
	"github.com/afrizalsebastian/go-gin-gorm/middleware"
	comment_services "github.com/afrizalsebastian/go-gin-gorm/services/comment"
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	claims, err := controllers.GetClaims(c)
	if err != nil {
		c.Error(err)
		return
	}

	postId, err := strconv.Atoi(c.Param("postId"))
	if err != nil {
		err := &middleware.CustomError{
			StatusCode: 400,
			Message:    "Path validation Error",
		}
		c.Error(err)
		return
	}

	var request dtos.CreateCommentRequest

	if err := c.ShouldBindBodyWithJSON(&request); err != nil {
		err := &middleware.CustomError{
			StatusCode: 400,
			Message:    "Validation Error",
		}
		c.Error(err)
		return
	}

	result, err := comment_services.Create(claims, postId, &request)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": true,
		"data":   result,
	})
}
