package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	UserID      uuid.UUID    `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"user_id"`
	Role        string       `gorm:"not null" json:"role"`
	CreatedAt   time.Time    `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time    `gorm:"autoUpdateTime" json:"updated_at"`

	ItemsLoaned []ItemLoaned `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;" json:"items_loaned"`

	AuthID uuid.UUID `gorm:"type:uuid" json:"auth_id"`
	Auth   Auth      `gorm:"foreignKey:AuthID;constraint:OnDelete:CASCADE;" json:"auth"`
}

// Fungsi untuk migrasi tabel User
func MigrateUser(db *gorm.DB) {
	db.AutoMigrate(&User{})
}
