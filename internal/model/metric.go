package model

import (
	"time"

	"github.com/google/uuid"
)

type Metric struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey;default:uuidv7()"`
	DeviceID    uuid.UUID `gorm:"type:uuid;index;not null"`
	WindowTitle string    `gorm:"not null"`
	ProcessName string    `gorm:"not null"`
	DurationMS  int       `gorm:"not null"`
	Timestamp   time.Time `gorm:"index;not null"`
}
