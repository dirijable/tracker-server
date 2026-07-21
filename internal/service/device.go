package service

import (
	"context"
	"fmt"
	"tracker-system/internal/dto"
	"tracker-system/internal/model"
	"tracker-system/internal/service/mapper"

	"github.com/google/uuid"
)

type DeviceRepository interface {
	FindAllByUserID(ctx context.Context, userID uuid.UUID) ([]model.Device, error)
	DeleteByID(ctx context.Context, id uuid.UUID) error
}

type DeviceService struct {
	repo DeviceRepository
}

func NewDeviceService(repo DeviceRepository) *DeviceService {
	return &DeviceService{
		repo: repo,
	}
}

func (s *DeviceService) FindAllByUserID(ctx context.Context, userID uuid.UUID) ([]dto.DeviceResponse, error) {
	modelDevices, err := s.repo.FindAllByUserID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("DeviceService.FindAllByUserID: %w", err)
	}
	devicesResponse := make([]dto.DeviceResponse, 0, len(modelDevices))
	for _, d := range modelDevices {
		devicesResponse = append(devicesResponse, mapper.DeviceModelToResponse(d))
	}
	return devicesResponse, nil
}

// TODO: добавить в хэндлер позже
func (s *DeviceService) DeleteByID(ctx context.Context, id uuid.UUID) error {
	if err := s.repo.DeleteByID(ctx, id); err != nil {
		return fmt.Errorf("DeviceService.DeleteByID: %w", err)
	}
	return nil
}
