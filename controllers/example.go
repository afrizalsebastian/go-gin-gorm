package controllers

import (
	"net/http"

	"github.com/afrizalsebastian/go-gin-gorm/middleware"
	"github.com/gin-gonic/gin"
)

func GetExample(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "/example routes",
	})
}

func GetExampleError(c *gin.Context) {
	err := &middleware.CustomError{
		StatusCode: 400,
		Message:    "Bad Request",
	}

	c.Error(err)
}
