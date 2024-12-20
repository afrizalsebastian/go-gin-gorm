package services

import (
	"net/http"

	"github.com/afrizalsebastian/go-gin-gorm/dtos"
	"github.com/afrizalsebastian/go-gin-gorm/middleware"
	"github.com/afrizalsebastian/go-gin-gorm/models"
	"github.com/afrizalsebastian/go-gin-gorm/repositories"
	"github.com/afrizalsebastian/go-gin-gorm/utils"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(createUserRequest *dtos.CreateUserRequest) (*models.User, error) {
	//Check email
	existEmail, err := repositories.GetUserByEmail(createUserRequest.Email)
	if err != nil {
		return nil, err
	}
	if existEmail != nil {
		return nil, middleware.NewCustomError(http.StatusBadRequest, "Bad Request. Email already used")
	}

	//Check username
	existUsername, err := repositories.GetUserByUsernae(createUserRequest.Username)
	if err != nil {
		return nil, err
	}
	if existUsername != nil {
		return nil, middleware.NewCustomError(http.StatusBadRequest, "Bad Request. Username already taken")
	}

	// Hash Password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(createUserRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Username: createUserRequest.Username,
		Password: string(hashedPassword),
		Email:    createUserRequest.Email,
	}

	return repositories.CreateUser(user)
}

func Login(loginRequest *dtos.LoginRequest) (string, error) {
	existUser, err := repositories.GetUserByEmail(loginRequest.Email)
	if err != nil {
		return "", err
	}
	if existUser == nil {
		return "", middleware.NewCustomError(http.StatusUnauthorized, "Email or Password incorrect.")
	}

	err = bcrypt.CompareHashAndPassword([]byte(existUser.Password), []byte(loginRequest.Password))
	if err != nil {
		return "", middleware.NewCustomError(http.StatusUnauthorized, "Email or Password incorrect.")
	}

	token, err := utils.CreateToken(existUser)
	if err != nil {
		return "", err
	}

	return token, nil
}
