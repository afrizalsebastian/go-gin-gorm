package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CustomError struct {
	StatusCode int
	Message    string
}

func (e *CustomError) Error() string {
	return e.Message
}

func ErrorHandling() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		if len(ctx.Errors) > 0 {
			err := ctx.Errors.Last().Err

			if customError, ok := err.(*CustomError); ok {
				ctx.JSON(customError.StatusCode, gin.H{
					"status": false,
					"error":  customError.Message,
				})
			} else {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"status": false,
					"error":  err.Error(),
				})
			}
		}
	}
}
