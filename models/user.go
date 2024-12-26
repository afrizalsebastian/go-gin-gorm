package models

import "time"

type Role string

const (
	AdminApp Role = "ADMIN"
	UserApp  Role = "USER"
)

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Username  string    `gorm:"unique;not null;type:varchar(100)" json:"username"`
	Email     string    `gorm:"unique;not null;type:varchar(100)" json:"email"`
	Password  string    `gorm:"not null;type:text" json:"-"`
	Role      Role      `gorm:"type:enum('USER','ADMIN');default:'USER'" json:"role"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	Profile   Profile   `gorm:"constraint:OnDelete:CASCADE;foreignKey:UserId;references:ID"`
}
