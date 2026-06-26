package handler

import (
	"context"
	"log"
	"net/http"
	"tracker-server/internal/dto"
	"tracker-server/internal/handler/decode"
	"tracker-server/internal/handler/encode"

	"github.com/google/uuid"
)

type MetricService interface {
	Save(ctx context.Context, deviceID uuid.UUID, request dto.CreateMetricRequest) (dto.MetricResponse, error)
}

type MetricHandler struct {
	s MetricService
}

func NewMetricHandler(s MetricService) *MetricHandler {
	return &MetricHandler{
		s: s,
	}
}

func (h *MetricHandler) Save(w http.ResponseWriter, r *http.Request) {
	var requestDto dto.CreateMetricRequest
	if err := decode.Decode(r, &requestDto); err != nil {
		//TODO err response
		log.Println(err)
		return
	}
	ctx := r.Context()
	//TODO мидлварь с получением айди девайса
	//deviceID, err := extract.ExtractUUID(ctx, "X-DEVICE-ID")
	deviceID, err := uuid.Parse("f47ac10b-58cc-4372-a567-0e02b2c3d479")
	if err != nil {
		log.Println(err)
		//TODO err response

		return
	}
	responseDto, err := h.s.Save(ctx, deviceID, requestDto)
	if err != nil {
		//TODO err resopnse
		log.Println(err)
		return
	}
	if err := encode.SendJSONResponse(w, http.StatusCreated, responseDto); err != nil {
		log.Println(err)
		return
	}
}
