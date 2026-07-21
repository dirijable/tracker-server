package service

import (
	"context"
	"fmt"
	"tracker-system/internal/dto"
	"tracker-system/internal/model"
	"tracker-system/internal/service/mapper"

	"github.com/google/uuid"
)

type MetricRepository interface {
	Save(ctx context.Context, metric model.Metric) (model.Metric, error)
}

type MetricService struct {
	repo MetricRepository
}

func NewMetricService(repo MetricRepository) *MetricService {
	return &MetricService{
		repo: repo,
	}
}

func (s *MetricService) Save(ctx context.Context, deviceID uuid.UUID, request dto.CreateMetricRequest) (dto.MetricResponse, error) {
	metric := mapper.MetricRequestToModel(request)
	metric.DeviceID = deviceID
	savedMetric, err := s.repo.Save(ctx, metric)
	if err != nil {
		return dto.MetricResponse{}, fmt.Errorf("service save metric: %w", err)
	}
	return mapper.MetricModelToResponse(savedMetric), err
}
