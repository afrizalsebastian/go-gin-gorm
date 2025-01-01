package repositories

import (
	"net/http"

	"github.com/afrizalsebastian/go-gin-gorm/config"
	"github.com/afrizalsebastian/go-gin-gorm/middleware"
	"github.com/afrizalsebastian/go-gin-gorm/models"
	"gorm.io/gorm"
)

func CreateComment(comment *models.Comment) error {
	return config.DB.Create(comment).Error
}

func GetCommentById(comment *models.Comment) error {
	if err := config.DB.Preload("User").Where("id = ? AND post_id = ?", comment.ID, comment.PostId).First(&comment).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return middleware.NewCustomError(http.StatusNotFound, "Comment not found")
		}
		return err
	}

	return nil
}
