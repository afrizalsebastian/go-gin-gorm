package controllers

import (
	"net/http"

	"github.com/afrizalsebastian/go-gin-gorm/middleware"
	"github.com/gin-gonic/gin"
)

func GetClaims(c *gin.Context) (*middleware.AppClaims, error) {
	claimsContext, exists := c.Get("user")
	if !exists {
		return nil, &middleware.CustomError{
			StatusCode: http.StatusUnauthorized,
			Message:    "403 Unathorized",
		}
	}

	claims, ok := claimsContext.(middleware.AppClaims)
	if !ok {
		return nil, &middleware.CustomError{
			StatusCode: http.StatusInternalServerError,
			Message:    "Something went wrong",
		}
	}

	return &claims, nil
}
