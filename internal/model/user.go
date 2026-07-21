package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;default:uuidv7()"`
	TGChatID  int64     `gorm:"type:bigint;not null;uniqueIndex"`
	CreatedAt time.Time `gorm:"not null"`
}
