package http

import (
	"context"
	"log"
	"net/http"
	"tracker-system/internal/dto"
	"tracker-system/internal/handler/http/decode"
	"tracker-system/internal/handler/http/encode"
)

type DeviceRegistrar interface {
	Register(ctx context.Context, request dto.RegisterDeviceRequest) (dto.RegisterDeviceResponse, error)
}

type DeviceRegistrarHandler struct {
	registrar DeviceRegistrar
}

func NewDeviceRegistrarHandler(registrar DeviceRegistrar) *DeviceRegistrarHandler {
	return &DeviceRegistrarHandler{
		registrar: registrar,
	}
}

func (h *DeviceRegistrarHandler) Register(w http.ResponseWriter, r *http.Request) {
	var regDto dto.RegisterDeviceRequest
	if err := decode.Decode(r, &regDto); err != nil {
		//TODO err response
		log.Println(err)
		return
	}
	token, err := h.registrar.Register(r.Context(), regDto)
	if err != nil {
		log.Println(err)
		return
	}

	if err := encode.SendJSONResponse(w, http.StatusCreated, token); err != nil {
		log.Println(err)
		return
	}
}
