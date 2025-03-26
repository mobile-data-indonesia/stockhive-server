package models

import (
	"time"

	"gorm.io/gorm"
)

type Vendor struct {
	VendorID      int            `gorm:"primaryKey;autoIncrement" json:"vendor_id"`
	VendorName    string         `gorm:"not null;unique" json:"vendor_name"`
	VendorDescription string     `gorm:"type:text" json:"vendor_description"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"` 
	CreatedAt       time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt       time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
}

