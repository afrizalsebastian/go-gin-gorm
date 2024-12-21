package repositories

import (
	"github.com/afrizalsebastian/go-gin-gorm/config"
	"github.com/afrizalsebastian/go-gin-gorm/models"
	"gorm.io/gorm"
)

func CreateUser(user *models.User) error {
	result := config.DB.Create(user)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	result := config.DB.Where("email = ?", email).First(&user)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error
	}
	return &user, nil
}

func GetUserByUsernae(username string) (*models.User, error) {
	var user models.User
	result := config.DB.Where("username = ?", username).First(&user)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error
	}
	return &user, nil
}
