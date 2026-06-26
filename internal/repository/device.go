package repository

import (
	"context"
	"errors"
	"fmt"
	"tracker-server/internal/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DeviceRepo struct {
	db *gorm.DB
}

func (r *DeviceRepo) Save(ctx context.Context, device model.Device) (model.Device, error) {
	if err := r.db.WithContext(ctx).Create(&device).Error; err != nil {
		return model.Device{}, fmt.Errorf("repo save device: %w", err)
	}
	return device, nil
}

func (r *DeviceRepo) FindByAPIToken(ctx context.Context, apiToken string) (model.Device, error) {
	var device model.Device
	if err := r.db.WithContext(ctx).Where("api_token = ?", apiToken).Take(&device).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.Device{}, err
		}
		return model.Device{}, fmt.Errorf("repo find by api token: %w", err)
	}
	return device, nil
}

func (r *DeviceRepo) FindAllByUserID(ctx context.Context, userID uuid.UUID) ([]model.Device, error) {
	var devices []model.Device
	if err := r.db.WithContext(ctx).Where(`user_id = ?`, userID).Find(&devices).Error; err != nil {
		return nil, fmt.Errorf("repo find all by userID: %w", err)
	}
	return devices, nil
}
