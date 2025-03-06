package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	UserID          uuid.UUID      `gorm:"type:uuid;primaryKey" json:"user_id"`
	Username        string         `gorm:"unique;not null" json:"username" binding:"required"`
	Password        string         `gorm:"not null" json:"password" binding:"required"`
	Email           string         `gorm:"unique" json:"email"`
	PhoneNumber     string         `json:"phone_number" binding:"required"`
	Title           string         `json:"title" binding:"required"`
	Role            string         `gorm:"type:varchar(10);not null;default:'staff'" json:"role" binding:"required,oneof=staff admin"`
	Department      string         `json:"department" binding:"required"`
	ProfileImageURL string         `json:"profile_image_url"` 
	LastLogin       *time.Time     `json:"last_login"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	CreatedAt       time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt       time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
}

type UserLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
    u.UserID = uuid.New()
    if u.ProfileImageURL == "" {
		u.ProfileImageURL = "default-profile.png" 
	}
    return nil
}