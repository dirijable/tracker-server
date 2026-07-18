package model

import (
	"time"

	"github.com/google/uuid"
)

type Device struct {
	ID        uuid.UUID `gorm:"not null;type:uuid;primaryKey;default:(-)"`
	UserID    uuid.UUID `gorm:"not null;type:uuid;index"`
	Name      string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null"`
}
