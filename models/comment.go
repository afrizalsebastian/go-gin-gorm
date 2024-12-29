package models

import "time"

type Comment struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Content   string    `gorm:"type:text" json:"content"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	UserId *uint `gorm:""`
	User   User  `gorm:"foreignKey:UserId;references:ID"`

	PostId *uint `gorm:"index"`
	Post   Post  `gorm:"foreignKey:PostId;references:ID"`
}
