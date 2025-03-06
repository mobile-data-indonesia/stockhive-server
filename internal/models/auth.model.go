package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Auth struct {
	AuthID   uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"auth_id"`
	Username string    `gorm:"unique;not null" json:"username"`
	Password string    `gorm:"not null" json:"password"`

	User *User `gorm:"foreignKey:AuthID;references:AuthID;constraint:OnDelete:CASCADE;"`
}

func MigrateAuth(db *gorm.DB) {
	db.AutoMigrate(&Auth{})
}
