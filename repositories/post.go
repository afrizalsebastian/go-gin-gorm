package repositories

import (
	"net/http"

	"github.com/afrizalsebastian/go-gin-gorm/config"
	"github.com/afrizalsebastian/go-gin-gorm/middleware"
	"github.com/afrizalsebastian/go-gin-gorm/models"
	"gorm.io/gorm"
)

func CreatePost(post *models.Post) error {
	return config.DB.Create(post).Error
}

func GetPostById(post *models.Post) error {
	result := config.DB.Preload("User.Profile").First(&post)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return middleware.NewCustomError(http.StatusNotFound, "Post not found")
		}
		return result.Error
	}

	return nil
}

func GetCountPost(count *int64) error {
	return config.DB.Model(&models.Post{}).Count(count).Error
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
