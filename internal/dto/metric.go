package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateMetricRequest struct {
	WindowTitle string    `json:"window_title" validate:"required"`
	ProcessName string    `json:"process_name" validate:"required"`
	DurationMS  int       `json:"duration_ms"  validate:"required,gt=0"`
	Timestamp   time.Time `json:"timestamp"    validate:"required"`
}

type MetricResponse struct {
	ID          uuid.UUID `json:"id"`
	DeviceID    uuid.UUID `json:"device_id"`
	WindowTitle string    `json:"window_title"`
	ProcessName string    `json:"process_name"`
	DurationMS  int       `json:"duration_ms"`
	Timestamp   time.Time `json:"timestamp"`
}
