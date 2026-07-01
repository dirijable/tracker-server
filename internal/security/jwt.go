package security

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

//var (
//	secretKey = []byte(os.Getenv("JWT_ACCESS_KEY"))
//)

type JWTManager struct {
	secretKey []byte
}

func NewJWTManager(secretKey []byte) *JWTManager {
	return &JWTManager{
		secretKey: secretKey,
	}
}

func (m *JWTManager) IssueAccessToken(deviceID uuid.UUID) (string, error) {
	//TODO: позже, возможно, можно будет сделать срок действия и рефреш токены
	claims := jwt.MapClaims{
		"d_id": deviceID.String(),
		"iat":  time.Now().Unix(),
		"iss":  "tracker-server",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(m.secretKey)
	if err != nil {
		return "", fmt.Errorf("issue access token: %w", err)
	}
	return tokenStr, nil
}

func (m *JWTManager) ParseToken(tokenStr string) (DeviceClaims, error) {
	deviceClaims := DeviceClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, &deviceClaims, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return m.secretKey, nil
	})

	if err != nil || !token.Valid {
		return DeviceClaims{}, fmt.Errorf("invalid token: %w", err)
	}

	return deviceClaims, nil
}
