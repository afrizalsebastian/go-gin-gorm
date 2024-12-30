package repositories

import (
	"github.com/afrizalsebastian/go-gin-gorm/config"
	"github.com/afrizalsebastian/go-gin-gorm/models"
)

func CreateComment(comment *models.Comment) error {
	return config.DB.Create(comment).Error
}
