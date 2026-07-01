package mapper

import (
	"time"
	"tracker-server/internal/domain"
	"tracker-server/internal/model"
)

func ActivationInfoToDeviceModel(info domain.ActivationInfo) model.Device {
	return model.Device{
		UserID:    info.UserID,
		Name:      info.DeviceName,
		CreatedAt: time.Now(),
	}
}
