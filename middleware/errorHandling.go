package middleware

import (
	"net/http"

	"github.com/afrizalsebastian/go-gin-gorm/utils"
	"github.com/gin-gonic/gin"
)

type CustomError struct {
	StatusCode int
	Message    string
}

func (e *CustomError) Error() string {
	return e.Message
}

func NewCustomError(statusCode int, message string) *CustomError {
	return &CustomError{
		StatusCode: statusCode,
		Message:    message,
	}
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
			} else if invalidToken, ok := err.(*utils.InvalidToken); ok {
				ctx.JSON(invalidToken.StatusCode, gin.H{
					"status": false,
					"error":  invalidToken.Message,
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
