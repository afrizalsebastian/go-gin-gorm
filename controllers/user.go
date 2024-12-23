package controllers

import (
	"net/http"

	"github.com/afrizalsebastian/go-gin-gorm/dtos"
	"github.com/afrizalsebastian/go-gin-gorm/middleware"
	user_services "github.com/afrizalsebastian/go-gin-gorm/services/user"
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

	user, err := user_services.CreateUser(&request)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": true,
		"data":   user,
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

	token, err := user_services.Login(&request)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data": gin.H{
			"token": token,
		},
	})
}

func DeleteUser(c *gin.Context) {
	claims, exists := c.Get("user")
	if !exists {
		c.Error(&middleware.CustomError{
			StatusCode: http.StatusUnauthorized,
			Message:    "403 Unathorized",
		})
		return
	}

	userId, ok := claims.(middleware.AppClaims)
	if !ok {
		c.Error(&middleware.CustomError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Something went wrong",
		})
		return
	}

	result, err := user_services.Delete(userId.ID)
	if err != nil {
		c.Error(&middleware.CustomError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   result,
	})
}
