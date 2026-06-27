package service

import (
	"context"
	"fmt"
	"tracker-server/internal/domain"
	"tracker-server/internal/model"
	"tracker-server/internal/service/mapper"

	"github.com/google/uuid"
)

type DeviceRepository interface {
	Save(ctx context.Context, device model.Device) (model.Device, error)
}

type RegistrarCache interface {
	Get(code string) (domain.ActivationInfo, error)
	DeleteIfMatching(code string, predicate func(info domain.ActivationInfo) bool) bool
}

type TokenIssuer interface {
	IssueAccessToken(deviceID uuid.UUID) (string, error)
}
type DeviceRegistrar struct {
	repo            DeviceRepository
	tokenIssuer     TokenIssuer
	activationCache RegistrarCache
}

func NewDeviceRegistrar(repository DeviceRepository, tokenIssuer TokenIssuer, ac RegistrarCache) *DeviceRegistrar {
	return &DeviceRegistrar{
		repo:            repository,
		tokenIssuer:     tokenIssuer,
		activationCache: ac,
	}
}

func (s *DeviceRegistrar) Register(ctx context.Context, code string) (string, error) {
	info, err := s.activationCache.Get(code)
	if err != nil {
		return "", fmt.Errorf("register device: %w", err)
	}
	deviceModel := mapper.ActivationInfoToDeviceModel(info)

	savedDevice, err := s.repo.Save(ctx, deviceModel);
	if err != nil {
		return "", fmt.Errorf("DeviceRegistrar.Register: %w", err)
	}

	token, err := s.tokenIssuer.IssueAccessToken(savedDevice.ID)
	if err != nil {
		return "", fmt.Errorf("tokenIssuer: %w", err)
	}

	s.activationCache.DeleteIfMatching(code, func(cachedInfo domain.ActivationInfo) bool {
		return cachedInfo.UserID == info.UserID
	})
	return token, nil
}
