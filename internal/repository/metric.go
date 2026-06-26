package repository

import (
	"context"
	"fmt"
	"tracker-server/internal/model"

	"gorm.io/gorm"
)

type MetricRepo struct {
	db *gorm.DB
}

func NewMetricRepo(db *gorm.DB) *MetricRepo {
	return &MetricRepo{
		db: db,
	}
}

func (m *MetricRepo) Save(ctx context.Context, metric model.Metric) (model.Metric, error) {
	if err := m.db.WithContext(ctx).Create(&metric).Error; err != nil {
		return model.Metric{}, fmt.Errorf("create metric: %w", err)
	}
	return metric, nil
}
