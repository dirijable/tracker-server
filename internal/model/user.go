package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primeryKey;default:(-)"`
	TGChatID  int64     `gorm:"type:bigint;not null;uniqueIndex"`
	CreatedAt time.Time `gorm:"not null"`
}
