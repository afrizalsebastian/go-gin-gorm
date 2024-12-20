package services

import (
	"github.com/afrizalsebastian/go-gin-gorm/dtos"
	"github.com/afrizalsebastian/go-gin-gorm/models"
)

func CreateUser(CreateUserRequest dtos.CreateUserRequest) models.User {

	user := models.User{
		Username: "Test",
		Password: "Test",
		Email:    "Test",
	}
	return user
}
