package handler

import (
	"context"
	"log"
	"net/http"
	"tracker-server/internal/dto"
	"tracker-server/internal/handler/encode"
	"tracker-server/internal/handler/extract"

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
	userID, err := extract.ExtractUUID(r.Context(), "X-USER-ID")
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
