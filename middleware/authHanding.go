package middleware

import (
	"net/http"
	"strings"

	"github.com/afrizalsebastian/go-gin-gorm/utils"
	"github.com/gin-gonic/gin"
)

func AuthenticationMiddleware(c *gin.Context) {
	header := c.GetHeader("Authorization")

	if header == "" || !strings.HasPrefix(header, "Bearer ") {
		c.Error(NewCustomError(http.StatusUnauthorized, "Token missing"))
		c.Abort()
		return
	}

	tokenString := strings.TrimPrefix(header, "Bearer ")

	decodedToken, err := utils.VerifyToken(tokenString)
	if err != nil {
		c.Error(err)
		c.Abort()
		return
	}

	claims := decodedToken.Claims

	c.Set("user", claims)
	c.Next()
}
