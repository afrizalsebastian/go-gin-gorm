package repositories

import (
	"github.com/afrizalsebastian/go-gin-gorm/config"
	"github.com/afrizalsebastian/go-gin-gorm/models"
	"gorm.io/gorm"
)

func CreatePost(post *models.Post) error {
	if err := config.DB.Create(post).Error; err != nil {
		return err
	}

	return nil
}

func GetPostById(postId int) (*models.Post, error) {
	var post models.Post
	result := config.DB.Preload("User.Profile").First(&post, postId)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error
	}

	return &post, nil
}

func UpdatePost(post *models.Post) error {
	return config.DB.Save(post).Error
}
