package repositories

import (
	"github.com/afrizalsebastian/go-gin-gorm/config"
	"github.com/afrizalsebastian/go-gin-gorm/models"
	"gorm.io/gorm"
)

func CreatePost(post *models.Post) error {
	return config.DB.Create(post).Error
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

func GetCountPost() (*int64, error) {
	var count int64
	if err := config.DB.Model(&models.Post{}).Count(&count).Error; err != nil {
		return nil, err
	}

	return &count, nil
}

func GetPost(rows int, offset int) ([]*models.Post, error) {
	var posts []*models.Post
	if err := config.DB.Limit(rows).Offset(offset).Preload("User.Profile").Find(&posts).Error; err != nil {
		return nil, err
	}

	return posts, nil
}

func UpdatePost(post *models.Post) error {
	return config.DB.Save(post).Error
}

func DeletePost(post *models.Post) error {
	return config.DB.Delete(post).Error
}
