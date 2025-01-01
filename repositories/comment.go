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
	query := map[string]interface{}{
		"id":      comment.ID,
		"post_id": comment.PostId,
	}
	if err := config.DB.Preload("User").Where(query).First(&comment).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return middleware.NewCustomError(http.StatusNotFound, "Comment not found")
		}
		return err
	}

	return nil
}

func GetCommentByIdAndUserId(comment *models.Comment) error {
	query := map[string]interface{}{
		"id":      comment.ID,
		"post_id": comment.PostId,
		"user_id": comment.UserId,
	}
	if err := config.DB.Preload("User").Where(query).First(&comment).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return middleware.NewCustomError(http.StatusNotFound, "Comment not found")
		}
		return err
	}

	return nil
}

func UpdateComment(comment *models.Comment) error {
	return config.DB.Save(&comment).Error
}
