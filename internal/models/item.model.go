package models

import (
	"time"

	"github.com/google/uuid"

	"gorm.io/gorm"
)


type Item struct {
	ItemID              uuid.UUID          `gorm:"type:uuid;primaryKey" json:"item_id"`
	ItemName            string             `gorm:"not null" json:"item_name"`
	VendorName          string             `json:"vendor_name"`
	ItemCategory          string            `json:"item_category"`
	ItemImageURL        string             `json:"item_image_url"`
	ItemLocationID      int                `gorm:"not null" json:"item_location_id"`
	ItemLocation        Location           `gorm:"foreignKey:ItemLocationID" json:"item_location"`
	HolderID            *string            `gorm:"type:uuid" json:"holder_id"` 
	Holder              *User              `gorm:"foreignKey:HolderID" json:"holder,omitempty"`
	ItemStatus 			string 				`gorm:"type:varchar(20);not null" json:"item_status" binding:"required,oneof=good pending new lost active damaged in_repair disposed"`
	DepreciationPeriod 	string 				`gorm:"type:varchar(10);not null" json:"depreciation_period" binding:"required,oneof=monthly yearly"`
	PurchaseDate        *time.Time         `json:"purchase_date,omitempty"`
	InitialValue        int                `json:"initial_value"`
	CurrentValue        int                `json:"current_value"`
	DepreciationRate    float64            `json:"depreciation_rate"`
	ItemDescription     string             `gorm:"type:text" json:"item_description"`
	DeletedAt           gorm.DeletedAt     `gorm:"index" json:"deleted_at,omitempty"`
	CreatedAt           time.Time          `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt           time.Time          `gorm:"autoUpdateTime" json:"updated_at"`
}

func (i *Item) BeforeCreate(tx *gorm.DB) (err error) {
    i.ItemID = uuid.New()
   
    return nil
}