package user_services

import (
	"net/http"

	"github.com/afrizalsebastian/go-gin-gorm/dtos"
	"github.com/afrizalsebastian/go-gin-gorm/middleware"
	"github.com/afrizalsebastian/go-gin-gorm/models"
	"github.com/afrizalsebastian/go-gin-gorm/repositories"
	"github.com/afrizalsebastian/go-gin-gorm/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func toUserResponse(user *models.User, profile *models.Profile) *dtos.UserResponse {
	var fullname string
	var bio string
	if profile != nil {
		fullname = profile.Fullname
		bio = profile.Bio
	}

	return &dtos.UserResponse{
		ID:       int(user.ID),
		Email:    user.Email,
		Username: user.Username,
		Role:     string(user.Role),
		Fullname: &fullname,
		Bio:      &bio,
	}
}

func Create(createUserRequest *dtos.CreateUserRequest) (*dtos.UserResponse, error) {
	//Check email
	existEmail, err := repositories.GetUserByEmail(createUserRequest.Email)
	if err != nil {
		return nil, err
	}
	if existEmail != nil {
		return nil, middleware.NewCustomError(http.StatusBadRequest, "Bad Request. Email already used")
	}

	//Check username
	existUsername, err := repositories.GetUserByUsername(createUserRequest.Username)
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
	err = repositories.CreateUser(user)
	if err != nil {
		return nil, err
	}

	profile := &models.Profile{
		UserId:   user.ID,
		Fullname: createUserRequest.Fullname,
		Bio:      createUserRequest.Bio,
	}
	err = repositories.CreateProfile(profile)
	if err != nil {
		return nil, err
	}

	return toUserResponse(user, profile), nil
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

func Get(userId int) (*dtos.UserResponse, error) {
	user, err := repositories.GetUserById(userId)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, middleware.NewCustomError(http.StatusNotFound, "This user not found.")
	}

	return toUserResponse(user, &user.Profile), nil
}

func Delete(id int) (*dtos.UserResponse, error) {
	deletedUser, err := repositories.DeleteUserById(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, middleware.NewCustomError(http.StatusNotFound, "This user not found.")

		}
		return nil, err
	}

	return toUserResponse(deletedUser, &deletedUser.Profile), nil
}
