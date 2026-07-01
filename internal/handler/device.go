package handler

import (
	"context"
	"log"
	"net/http"
	"tracker-server/internal/dto"
	"tracker-server/internal/handler/decode"
	"tracker-server/internal/handler/encode"
)

type DeviceRegistrar interface {
	Register(ctx context.Context, request dto.RegisterDeviceRequest) (dto.RegisterDeviceResponse, error)
}

type DeviceHandler struct {
	deviceRegistrar DeviceRegistrar
}

func NewDeviceHandler(r DeviceRegistrar) *DeviceHandler {
	return &DeviceHandler{
		deviceRegistrar: r,
	}
}

func (h *DeviceHandler) Register(w http.ResponseWriter, r *http.Request) {
	var regDto dto.RegisterDeviceRequest
	if err := decode.Decode(r, &regDto); err != nil {
		//TODO err response
		log.Println(err)
		return
	}
	token, err := h.deviceRegistrar.Register(r.Context(), regDto)
	if err != nil {
		log.Println(err)
		return
	}

	if err := encode.SendJSONResponse(w, http.StatusCreated, token); err != nil {
		log.Println(err)
		return
	}
}
