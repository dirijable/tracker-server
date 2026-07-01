package dto

type RegisterDeviceRequest struct {
	Code string `json:"code" validate:"required"`
}

type RegisterDeviceResponse struct {
	Token string `json:"token"`
}
