package domain

import (
	"time"

	"github.com/google/uuid"
)

type ActivationInfo struct {
	UserID     uuid.UUID
	DeviceName string
	ExpireAt   time.Time
}
