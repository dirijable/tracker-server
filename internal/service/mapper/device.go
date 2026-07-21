package mapper

import (
	"time"
	"tracker-system/internal/domain"
	"tracker-system/internal/dto"
	"tracker-system/internal/model"
)

func ActivationInfoToDeviceModel(info domain.ActivationInfo) model.Device {
	return model.Device{
		UserID:    info.UserID,
		Name:      info.DeviceName,
		CreatedAt: time.Now(),
	}
}

func DeviceModelToResponse(deviceModel model.Device) dto.DeviceResponse {
	return dto.DeviceResponse{
		ID:        deviceModel.ID,
		Name:      deviceModel.Name,
		CreatedAt: deviceModel.CreatedAt,
	}
}
