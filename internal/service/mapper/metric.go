package mapper

import (
	"tracker-server/internal/dto"
	"tracker-server/internal/model"
)

func RequestToModel(request dto.CreateMetricRequest) model.Metric {
	return model.Metric{
		WindowTitle: request.WindowTitle,
		ProcessName: request.ProcessName,
		DurationMS:  request.DurationMS,
		Timestamp:   request.Timestamp,
	}
}

func ModelToResponse(metric model.Metric) dto.MetricResponse {
	return dto.MetricResponse{
		ID:          metric.ID,
		DeviceID:    metric.DeviceID,
		WindowTitle: metric.WindowTitle,
		ProcessName: metric.ProcessName,
		DurationMS:  metric.DurationMS,
		Timestamp:   metric.Timestamp,
	}
}
