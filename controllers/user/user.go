package user_controllers

import (
	"net/http"

	"github.com/afrizalsebastian/go-gin-gorm/dtos"
	"github.com/afrizalsebastian/go-gin-gorm/middleware"
	user_services "github.com/afrizalsebastian/go-gin-gorm/services/user"
	"github.com/gin-gonic/gin"
)

func getClaims(c *gin.Context) (*middleware.AppClaims, error) {
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

func Create(c *gin.Context) {
	var request dtos.CreateUserRequest

	if err := c.ShouldBindBodyWithJSON(&request); err != nil {
		err := &middleware.CustomError{
			StatusCode: 400,
			Message:    "Validation Error",
		}
		c.Error(err)
		return
	}

	user, err := user_services.Create(&request)
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

func Get(c *gin.Context) {
	claims, err := getClaims(c)
	if err != nil {
		c.Error(err)
		return
	}

	result, err := user_services.Get(claims.ID)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   result,
	})
}

func Update(c *gin.Context) {
	claims, err := getClaims(c)
	if err != nil {
		c.Error(err)
		return
	}

	var request dtos.UpdateUserRequest
	if err := c.ShouldBindBodyWithJSON(&request); err != nil {
		err := &middleware.CustomError{
			StatusCode: 400,
			Message:    "Validation Error",
		}
		c.Error(err)
		return
	}

	result, err := user_services.Update(claims.ID, &request)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   result,
	})
}

func Delete(c *gin.Context) {
	claims, err := getClaims(c)
	if err != nil {
		c.Error(err)
		return
	}

	result, err := user_services.Delete(claims.ID)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   result,
	})
}
