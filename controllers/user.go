package controllers

import (
	"net/http"

	"github.com/afrizalsebastian/go-gin-gorm/dtos"
	"github.com/afrizalsebastian/go-gin-gorm/middleware"
	"github.com/afrizalsebastian/go-gin-gorm/services"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var request dtos.CreateUserRequest

	if err := c.ShouldBindBodyWithJSON(&request); err != nil {
		err := &middleware.CustomError{
			StatusCode: 400,
			Message:    "Validation Error",
		}
		c.Error(err)
		return
	}

	user := services.CreateUser(request)

	c.JSON(http.StatusCreated, gin.H{
		"status": true,
		"data": gin.H{
			"username": user.Username,
		},
	})
}
