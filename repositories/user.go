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

func GetUserById(id int) (*models.User, error) {
	var user models.User
	result := config.DB.First(&user, id)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error
	}

	return &user, nil
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

func GetUserByUsername(username string) (*models.User, error) {
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

func DeleteUserById(id int) (*models.User, error) {
	var user models.User
	if err := config.DB.Preload("Profile").First(&user, id).Error; err != nil {
		return nil, err
	}

	if err := config.DB.Delete(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
