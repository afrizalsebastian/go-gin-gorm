package middleware

import (
	"net/http"
	"strings"

	"github.com/afrizalsebastian/go-gin-gorm/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type AppClaims struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

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

	claims := decodedToken.Claims.(jwt.MapClaims)
	appClaims := AppClaims{
		ID:       int(claims["id"].(float64)),
		Username: claims["username"].(string),
		Email:    claims["email"].(string),
		Role:     claims["role"].(string),
	}

	c.Set("user", appClaims)
	c.Next()
}
