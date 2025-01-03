package post_controllers

import (
	"net/http"
	"strconv"

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
	c.JSON(http.StatusCreated, gin.H{
		"status": true,
		"data":   result,
	})
}

func Get(c *gin.Context) {
	rowsQuery := c.DefaultQuery("rows", "5")
	pageQuery := c.DefaultQuery("page", "1")

	rows, err := strconv.Atoi(rowsQuery)
	if err != nil {
		err := &middleware.CustomError{
			StatusCode: 400,
			Message:    "Rows validation Error",
		}
		c.Error(err)
		return
	}

	page, err := strconv.Atoi(pageQuery)
	if err != nil {
		err := &middleware.CustomError{
			StatusCode: 400,
			Message:    "Rows validation Error",
		}
		c.Error(err)
		return
	}

	result, err := post_services.Get(rows, page)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   result,
	})
}

func GetById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("postId"))
	if err != nil {
		err := &middleware.CustomError{
			StatusCode: 400,
			Message:    "Path validation Error",
		}
		c.Error(err)
		return
	}

	result, err := post_services.GetById(id)
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
	claims, err := controllers.GetClaims(c)
	if err != nil {
		c.Error(err)
		return
	}

	var request dtos.UpdatePostRequest
	if err := c.ShouldBindBodyWithJSON(&request); err != nil {
		err := &middleware.CustomError{
			StatusCode: 400,
			Message:    "Validation Error",
		}
		c.Error(err)
		return
	}

	id, err := strconv.Atoi(c.Param("postId"))
	if err != nil {
		err := &middleware.CustomError{
			StatusCode: 400,
			Message:    "Path validation Error",
		}
		c.Error(err)
		return
	}

	result, err := post_services.Update(claims, id, &request)
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
	claims, err := controllers.GetClaims(c)
	if err != nil {
		c.Error(err)
		return
	}

	id, err := strconv.Atoi(c.Param("postId"))
	if err != nil {
		err := &middleware.CustomError{
			StatusCode: 400,
			Message:    "Path validation Error",
		}
		c.Error(err)
		return
	}

	result, err := post_services.Delete(claims, id)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   result,
	})
}
