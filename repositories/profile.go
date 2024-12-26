package repositories

import (
	"github.com/afrizalsebastian/go-gin-gorm/config"
	"github.com/afrizalsebastian/go-gin-gorm/models"
)

func CreateProfile(profile *models.Profile) error {
	if err := config.DB.Create(profile).Error; err != nil {
		return err
	}

	return nil
}

func UpdateProfile(profile *models.Profile) error {
	if err := config.DB.Save(profile).Error; err != nil {
		return err
	}

	return nil
}

func DeleteProfile(id int) (*models.Profile, error) {
	var profile models.Profile
	result := config.DB.Delete(&profile, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &profile, nil
}
