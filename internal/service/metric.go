package service

import (
	"context"
	"fmt"
	"tracker-server/internal/dto"
	"tracker-server/internal/model"
	"tracker-server/internal/service/mapper"

	"github.com/google/uuid"
)

type MetricRepository interface {
	Save(ctx context.Context, metric model.Metric) (model.Metric, error)
}

type MetricService struct {
	repo MetricRepository
}

func (s *MetricService) Save(ctx context.Context, deviceID uuid.UUID, request dto.CreateMetricRequest) (dto.MetricResponse, error) {
	metric := mapper.RequestToModel(request)
	metric.DeviceID = deviceID
	savedMetric, err := s.repo.Save(ctx, metric)
	if err != nil {
		return dto.MetricResponse{}, fmt.Errorf("service save metric: %w", err)
	}
	return mapper.ModelToResponse(savedMetric), err
}
