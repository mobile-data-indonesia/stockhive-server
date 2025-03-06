package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Item struct {
	ItemID   uuid.UUID `gorm:"type:uuid;primaryKey" json:"auth_id"`
	ItemName string    `gorm:"not null" json:"item_name"`
	ItemDesc string    `gorm:"not null" json:"item_desc"`
	ItemPrice int    `gorm:"not null" json:"item_price"`
	ItemStatus string    `gorm:"not null" json:"item_status"`
	ItemImage string    `gorm:"not null" json:"item_image"` 
	CreatedAt   time.Time    `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time    `gorm:"autoUpdateTime" json:"updated_at"`
}

func MigrateItem(db *gorm.DB){
	db.AutoMigrate(&Item{})
}