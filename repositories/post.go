package repositories

import (
	"github.com/afrizalsebastian/go-gin-gorm/config"
	"github.com/afrizalsebastian/go-gin-gorm/models"
)

func CreatePost(post *models.Post) error {
	if err := config.DB.Create(post).Error; err != nil {
		return err
	}

	return nil
}
