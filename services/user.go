package services

import (
	"github.com/afrizalsebastian/go-gin-gorm/dtos"
	"github.com/afrizalsebastian/go-gin-gorm/middleware"
	"github.com/afrizalsebastian/go-gin-gorm/models"
	"github.com/afrizalsebastian/go-gin-gorm/repositories"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(CreateUserRequest *dtos.CreateUserRequest) (*models.User, error) {

	//Check email
	existEmail, err := repositories.GetUserByEmail(CreateUserRequest.Email)
	if existEmail != nil {
		return nil, middleware.NewCustomError(400, "Bad Request. Email already used")
	}
	if err != nil {
		return nil, err
	}

	//Check username
	existUsername, err := repositories.GetUserByUsernae(CreateUserRequest.Username)
	if existUsername != nil {
		return nil, middleware.NewCustomError(400, "Bad Request. Username already taken")
	}
	if err != nil {
		return nil, err
	}

	// Hash Password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(CreateUserRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Username: CreateUserRequest.Username,
		Password: string(hashedPassword),
		Email:    CreateUserRequest.Email,
	}

	return repositories.CreateUser(user)
}
