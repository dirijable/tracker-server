package repository

import (
	"context"
	"fmt"
	"tracker-system/internal/model"

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

func (r *DeviceRepo) FindAllByUserID(ctx context.Context, userID uuid.UUID) ([]model.Device, error) {
	var devices []model.Device
	if err := r.db.WithContext(ctx).Where(`user_id = ?`, userID).Find(&devices).Error; err != nil {
		return nil, fmt.Errorf("repo find all by userID: %w", err)
	}
	return devices, nil
}

func (r *DeviceRepo) DeleteByID(ctx context.Context, id uuid.UUID) error {
	sql := `DELETE FROM devices
			WHERE id = ?`
	if err := r.db.WithContext(ctx).Raw(sql, id).Error; err != nil {
		return fmt.Errorf("repo delete by id: %w", err)
	}
	return nil
}
