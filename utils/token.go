package utils

import (
	"net/http"
	"os"
	"time"

	"github.com/afrizalsebastian/go-gin-gorm/models"
	"github.com/golang-jwt/jwt/v5"
)

type JwtClaims struct {
	jwt.RegisteredClaims
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

var JWT_SECRET = []byte(os.Getenv("JWT_KEY"))

type InvalidToken struct {
	StatusCode int
	Message    string
}

func (e *InvalidToken) Error() string {
	return e.Message
}

func CreateToken(user *models.User) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       user.ID,
		"username": user.Username,
		"email":    user.Email,
		"role":     user.Role,
		"exp":      time.Now().Add(6 * time.Hour).Unix(),
		"iat":      time.Now().Unix(),
		"iss":      "localhost:8000",
		"aud":      "localhost:8000",
		"sub":      user.Username,
	})

	token, err := claims.SignedString(JWT_SECRET)
	if err != nil {
		return "", err
	}

	return token, nil
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	claims, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return JWT_SECRET, nil
	})

	if err != nil {
		return nil, err
	}

	if !claims.Valid {
		return nil, &InvalidToken{
			Message:    "Invalid Token",
			StatusCode: http.StatusUnauthorized,
		}
	}

	return claims, nil
}
