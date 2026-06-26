package service

import (
	"context"
	"fmt"
	"time"
	"tracker-server/internal/domain"
	"tracker-server/internal/model"
)

type DeviceRepository interface {
	Save(ctx context.Context, device model.Device) (model.Device, error)
}

type RegistrarCache interface {
	Get(code string) (domain.ActivationInfo, error)
	DeleteIfMatching(code string, predicate func(info domain.ActivationInfo) bool) bool
}

type TokenGenerator interface {
	GenerateAPITokenHEX() (string, error)
	HashTokenSHA256(token string) string
}

type DeviceRegistrar struct {
	repo            DeviceRepository
	tokenManager    TokenGenerator
	activationCache RegistrarCache
}

func NewDeviceRegistrar(repository DeviceRepository, tokenManager TokenGenerator, ac RegistrarCache) *DeviceRegistrar {
	return &DeviceRegistrar{
		repo:            repository,
		tokenManager:    tokenManager,
		activationCache: ac,
	}
}

func (s *DeviceRegistrar) Register(ctx context.Context, code string) (string, error) {
	info, err := s.activationCache.Get(code)
	if err != nil {
		return "", fmt.Errorf("register device: %w", err)
	}
	hexToken, err := s.tokenManager.GenerateAPITokenHEX()
	if err != nil {
		return "", fmt.Errorf("generate token hex: %w", err)
	}
	hashedToken := s.tokenManager.HashTokenSHA256(hexToken)
	deviceModel := model.Device{
		UserID:    info.UserID,
		Name:      info.DeviceName,
		ApiToken:  hashedToken,
		CreatedAt: time.Now(),
	}
	if _, err := s.repo.Save(ctx, deviceModel); err != nil {
		return "", fmt.Errorf("DeviceRegistrar.Register: %w", err)
	}

	s.activationCache.DeleteIfMatching(code, func(cachedInfo domain.ActivationInfo) bool {
		return cachedInfo.UserID == info.UserID
	}) //TODO возможно как-то обработать, что кто-то уже удалил запись
	return hexToken, nil
}
