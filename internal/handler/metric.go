package handler

import (
	"context"
	"net/http"
	"tracker-server/internal/dto"
	"tracker-server/internal/handler/decode"
	"tracker-server/internal/handler/encode"
	"tracker-server/internal/handler/extract"

	"github.com/google/uuid"
)

type MetricService interface {
	Save(ctx context.Context, deviceID uuid.UUID, request dto.CreateMetricRequest) (dto.MetricResponse, error)
}

type MetricHandler struct {
	s MetricService
}

func (h *MetricHandler) Save(w http.ResponseWriter, r *http.Request) {
	var requestDto dto.CreateMetricRequest
	if err := decode.Decode(r, &requestDto); err != nil {
		//TODO err response
		return
	}
	ctx := r.Context()
	//TODO мидлварь с получением айди девайса
	deviceID, err := extract.ExtractUUID(ctx, "X-DEVICE-ID")
	if err != nil {
		//TODO err response

		return
	}
	responseDto, err := h.s.Save(ctx, deviceID, requestDto)
	if err != nil {
		//TODO err resopnse
		return
	}
	if err := encode.SendJSONResponse(w, http.StatusCreated, responseDto); err != nil {
		return
	}
}
