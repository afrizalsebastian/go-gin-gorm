package models

import (
	"time"
)

type Post struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Title     string    `gorm:"type:varchar(250);not null" json:"title"`
	Content   string    `gorm:"type:text" json:"content"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	UserId   *uint     `gorm:"index"`
	User     *User     `gorm:"foreignKey:UserId;references:ID"`
	Comments []Comment `gorm:"foreignKey:PostId;references:ID;constraint:OnDelete:SET NULL"`
}
