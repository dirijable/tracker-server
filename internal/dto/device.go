package dto

import (
	"time"

	"github.com/google/uuid"
)

type RegisterDeviceRequest struct {
	Code string `json:"code" validate:"required"`
}

type RegisterDeviceResponse struct {
	Token string `json:"token"`
}

type DeviceResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}
