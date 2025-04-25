package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ItemRequest struct {
	ItemRequestID   uuid.UUID      `gorm:"type:uuid;primaryKey" json:"item_request_id"`
	RequesterID     uuid.UUID      `gorm:"type:uuid" json:"requester_id"`
	Requester       User           `gorm:"foreignKey:RequesterID" json:"requester"`
	ItemName        string         `gorm:"not null" json:"item_name"`
	RequestVendor   string         `gorm:"not null" json:"request_vendor"`
	ItemCategory    string         `json:"item_category"`
	RequestQuantity int            `gorm:"not null" json:"item_quantity"`
	RequestNotes    string         `gorm:"type:text" json:"request_notes"`
	RequestStatus   string         `gorm:"type:varchar(20);not null" json:"request_status" binding:"required,oneof=pending approved rejected"`
	CreatedAt       time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt       time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}
