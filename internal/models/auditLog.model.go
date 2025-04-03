package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AuditLog struct {
	AuditID      	uuid.UUID       `gorm:"type:uuid;primaryKey" json:"audit_id"`
	AuditorID 		string 			`gorm:"type:uuid;not null" json:"auditor_id"`
	Auditor 		User 			`gorm:"foreignKey:AuditorID" json:"auditor" binding:"-"`
	AuditStatus 	string    		`gorm:"type:text" json:"audit_status"`
	AuditNotes 		string    		`gorm:"type:text" json:"audit_notes"`
	DeletedAt       gorm.DeletedAt 	`gorm:"index" json:"deleted_at,omitempty"` 
	AuditDate       time.Time      	`gorm:"autoCreateTime" json:"audit_date"`
}

func (i *AuditLog) BeforeCreate(tx *gorm.DB) (err error) {
    i.AuditID = uuid.New()
   
    return nil
}