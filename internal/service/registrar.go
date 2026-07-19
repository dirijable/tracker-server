package service

import (
	"context"
	"fmt"
	"tracker-server/internal/domain"
	"tracker-server/internal/dto"
	"tracker-server/internal/model"
	"tracker-server/internal/service/mapper"

	"github.com/google/uuid"
)

type DeviceSaver interface {
	Save(ctx context.Context, device model.Device) (model.Device, error)
	DeleteByID(ctx context.Context, id uuid.UUID) error
}

type RegistrarCache interface {
	Get(code string) (domain.ActivationInfo, error)
	DeleteIfMatching(code string, predicate func(info domain.ActivationInfo) bool) bool
}

type TokenIssuer interface {
	IssueAccessToken(deviceID uuid.UUID) (string, error)
}

type DeviceRegistrar struct {
	repo            DeviceSaver
	tokenIssuer     TokenIssuer
	activationCache RegistrarCache
}

func NewDeviceRegistrar(repository DeviceSaver, tokenIssuer TokenIssuer, ac RegistrarCache) *DeviceRegistrar {
	return &DeviceRegistrar{
		repo:            repository,
		tokenIssuer:     tokenIssuer,
		activationCache: ac,
	}
}

func (s *DeviceRegistrar) Register(ctx context.Context, request dto.RegisterDeviceRequest) (dto.RegisterDeviceResponse, error) {
	info, err := s.activationCache.Get(request.Code)
	if err != nil {
		return dto.RegisterDeviceResponse{}, fmt.Errorf("register device: %w", err)
	}
	deviceModel := mapper.ActivationInfoToDeviceModel(info)

	savedDevice, err := s.repo.Save(ctx, deviceModel)
	if err != nil {
		return dto.RegisterDeviceResponse{}, fmt.Errorf("DeviceRegistrar.Register: %w", err)
	}

	token, err := s.tokenIssuer.IssueAccessToken(savedDevice.ID)
	if err != nil {
		if err := s.repo.DeleteByID(ctx, savedDevice.ID); err != nil {
			return dto.RegisterDeviceResponse{}, fmt.Errorf("delete device by id after fail to issue token: %w", err)
		}
		return dto.RegisterDeviceResponse{}, fmt.Errorf("tokenIssuer: %w", err)
	}

	s.activationCache.DeleteIfMatching(request.Code, func(cachedInfo domain.ActivationInfo) bool {
		return cachedInfo.UserID == info.UserID
	})
	return dto.RegisterDeviceResponse{Token: token}, nil
}
