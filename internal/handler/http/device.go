package http

import (
	"context"
	"log"
	"net/http"
	"tracker-system/internal/dto"
	"tracker-system/internal/handler/http/encode"
	"tracker-system/internal/handler/http/extract"

	"github.com/google/uuid"
)

type DeviceManager interface {
	FindAllByUserID(ctx context.Context, userID uuid.UUID) ([]dto.DeviceResponse, error)
}

type DeviceHandler struct {
	svc DeviceManager
}

func NewDeviceHandler(svc DeviceManager) *DeviceHandler {
	return &DeviceHandler{
		svc: svc,
	}
}

func (h *DeviceHandler) FindAllByUserID(w http.ResponseWriter, r *http.Request) {
	userID, err := extract.UUIDFromCtx(r.Context(), "X-USER-ID")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	response, err := h.svc.FindAllByUserID(r.Context(), userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if err := encode.SendJSONResponse(w, http.StatusOK, response); err != nil {
		log.Println(err)
	}
}
