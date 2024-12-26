package models

import "time"

type Profile struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Fullname  string    `gorm:"type:varchar(250);not null" json:"fullname"`
	Bio       string    `gorm:"type:text" json:"bio"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	UserId    uint      `gorm:"not null;"`
}
