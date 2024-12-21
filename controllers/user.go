package controllers

import (
	"net/http"

	"github.com/afrizalsebastian/go-gin-gorm/dtos"
	"github.com/afrizalsebastian/go-gin-gorm/middleware"
	"github.com/afrizalsebastian/go-gin-gorm/models"
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

	user, err := services.CreateUser(&request)
	if err != nil {
		c.Error(err)
		return
	}

	userResponse := map[string]interface{}{
		"username": (*user)["username"].(string),
		"id":       (*user)["id"].(uint),
		"email":    (*user)["email"].(string),
		"role":     ((*user)["role"].(models.Role)),
		"fullname": (*user)["fullname"].(string),
		"bio":      (*user)["bio"].(string),
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": true,
		"data":   userResponse,
	})
}

func Login(c *gin.Context) {
	var request dtos.LoginRequest

	if err := c.ShouldBindBodyWithJSON(&request); err != nil {
		err := &middleware.CustomError{
			StatusCode: 400,
			Message:    "Validation Error",
		}
		c.Error(err)
		return
	}

	token, err := services.Login(&request)
	if err != nil {
		c.Error(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data": gin.H{
			"token": token,
		},
	})
}
