package post_controllers

import (
	"net/http"

	"github.com/afrizalsebastian/go-gin-gorm/controllers"
	"github.com/afrizalsebastian/go-gin-gorm/dtos"
	"github.com/afrizalsebastian/go-gin-gorm/middleware"
	post_services "github.com/afrizalsebastian/go-gin-gorm/services/post"
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	claims, err := controllers.GetClaims(c)
	if err != nil {
		c.Error(err)
		return
	}

	var request dtos.CreatePostRequest

	if err := c.ShouldBindBodyWithJSON(&request); err != nil {
		err := &middleware.CustomError{
			StatusCode: 400,
			Message:    "Validation Error",
		}
		c.Error(err)
		return
	}

	result, err := post_services.Create(claims.ID, &request)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   result,
	})
}