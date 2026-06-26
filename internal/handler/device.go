package handler

import (
	"context"
	"net/http"
)

type DeviceRegistrar interface {
	Register(ctx context.Context, code string) (string, error)
}

type DeviceHandler struct {
	r DeviceRegistrar
}

func NewDeviceHandler(r DeviceRegistrar) *DeviceHandler {
	return &DeviceHandler{
		r: r,
	}
}

func (h *DeviceHandler) Register(w http.ResponseWriter, r *http.Request) {

}
