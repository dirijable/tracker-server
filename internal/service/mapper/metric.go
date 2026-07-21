package mapper

import (
	"tracker-system/internal/dto"
	"tracker-system/internal/model"
)

func MetricRequestToModel(request dto.CreateMetricRequest) model.Metric {
	return model.Metric{
		WindowTitle: request.WindowTitle,
		ProcessName: request.ProcessName,
		DurationMS:  request.DurationMS,
		Timestamp:   request.Timestamp,
	}
}

func MetricModelToResponse(metric model.Metric) dto.MetricResponse {
	return dto.MetricResponse{
		ID:          metric.ID,
		DeviceID:    metric.DeviceID,
		WindowTitle: metric.WindowTitle,
		ProcessName: metric.ProcessName,
		DurationMS:  metric.DurationMS,
		Timestamp:   metric.Timestamp,
	}
}
