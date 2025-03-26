package models

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	CategoryID      int            `gorm:"primaryKey;autoIncrement" json:"category_id"`
	CategoryName    string         `gorm:"not null;unique" json:"category_name"`
	CategoryDescription string     `gorm:"type:text" json:"category_description"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"` 
	CreatedAt       time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt       time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
}

