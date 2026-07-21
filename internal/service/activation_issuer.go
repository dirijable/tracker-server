package service

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
	"tracker-system/internal/domain"

	"github.com/google/uuid"
)

type IssuerCache interface {
	PutIfAbsent(code string, info domain.ActivationInfo) bool
}

type ActivationIssuer struct {
	activationCache IssuerCache
	codeTTL         time.Duration
	codeLength      int
}

func NewActivationIssuer(cache IssuerCache, ttl time.Duration, length int) *ActivationIssuer {
	return &ActivationIssuer{
		activationCache: cache,
		codeTTL:         ttl,
		codeLength:      length,
	}
}

func (s *ActivationIssuer) IssueCode(ctx context.Context, userID uuid.UUID, deviceName string) (string, error) {
	activationInfo := domain.ActivationInfo{
		UserID:     userID,
		DeviceName: deviceName,
		ExpireAt:   time.Now().Add(s.codeTTL),
	}
	for {
		code, err := generateActivationCode(s.codeLength)
		if err != nil {
			return "", fmt.Errorf("generate code: %w", err)
		}
		if s.activationCache.PutIfAbsent(code, activationInfo) {
			return code, nil
		}
	}
}

func generateActivationCode(length int) (string, error) {
	const telegramCharset = "23456789ABCDEFGHJKLMNPQRSTUVWXYZ"
	codeSymbols := make([]byte, length)
	maxLength := big.NewInt(int64(len(telegramCharset)))
	for i := range length {
		idx, err := rand.Int(rand.Reader, maxLength)
		if err != nil {
			return "", fmt.Errorf("generate activation code: %w", err)
		}
		codeSymbols[i] = telegramCharset[idx.Int64()]
	}
	return string(codeSymbols), nil
}
