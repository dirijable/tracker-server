package security

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type DeviceClaims struct {
	DeviceID uuid.UUID `json:"d_id"`
	jwt.RegisteredClaims
}
