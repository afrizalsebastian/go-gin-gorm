package utils

import (
	"os"
	"time"

	"github.com/afrizalsebastian/go-gin-gorm/models"
	"github.com/golang-jwt/jwt/v5"
)

type JwtClaims struct {
	ID       int `json:"id"`
	Username int `json:"username"`
	Email    int `json:"email"`
	Role     int `json:"role"`
}

var JWT_SECRET = []byte(os.Getenv("JWT_KEY"))

func CreateToken(user *models.User) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       user.ID,
		"username": user.Username,
		"email":    user.Email,
		"role":     user.Role,
		"exp":      time.Now().Add(2 * time.Minute).Unix(),
		"iat":      time.Now().Unix(),
	})

	token, err := claims.SignedString(JWT_SECRET)
	if err != nil {
		return "", err
	}

	return token, nil
}
