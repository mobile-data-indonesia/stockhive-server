package models

import (
	"time"

	"github.com/google/uuid"
)

type ActivityLog struct {
	ActivityID          uint      `gorm:"unique;primaryKey;autoIncrement" json:"activity_id"`
	ActorID             uuid.UUID `gorm:"type:uuid;not null" json:"actor_id"`
	Actor               User      `gorm:"foreignKey:ActorID" json:"actor" binding:"-"`
	Action              string    `gorm:"type:text" json:"action"`
	ActivityDescription string    `gorm:"type:text" json:"activity_description"`
	ActivityDate        time.Time `gorm:"autoCreateTime" json:"activity_date"`
}
