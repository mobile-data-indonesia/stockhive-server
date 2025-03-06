package models

import (
	"time"

	"gorm.io/gorm"
)

type Location struct {
	LocationID      int            `gorm:"primaryKey;autoIncrement" json:"location_id"`
	LocationName    string         `gorm:"not null;unique" json:"location_name"`
	LocationDescription string     `gorm:"type:text" json:"location_description"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"` 
	CreatedAt       time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt       time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
}

