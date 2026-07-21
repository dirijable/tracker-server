package bot

import (
	"context"
	"fmt"
	"tracker-system/internal/handler/bot/extract"
	"tracker-system/pkg/ctxx"

	"github.com/google/uuid"
	"gopkg.in/telebot.v3"
)

type ActivationCodeIssuer interface {
	IssueCode(ctx context.Context, userID uuid.UUID, deviceName string) (string, error)
}

type DeviceHandler struct {
	aci ActivationCodeIssuer
}

func NewDeviceHandler(aci ActivationCodeIssuer) *DeviceHandler {
	return &DeviceHandler{
		aci: aci,
	}
}

func (h *DeviceHandler) AddDevice(c telebot.Context) error {
	ctx := ctxx.From(c)
	deviceName := c.Message().Payload
	if deviceName == "" {
		return c.Send("usage: /add_device Device Name")
	}
	userID, err := extract.UUIDFromTeleCtx(c, "userID")
	if err != nil {
		return c.Send("user's id is not send")
	}
	code, err := h.aci.IssueCode(ctx, userID, deviceName)
	if err != nil {
		return c.Send("Failed to generate the code. Please try again later")
	}
	msg := fmt.Sprintf("Your access code: <code>%s</code>\n\nEnter this code in Tracker program on your device", code)
	return c.Send(msg, telebot.ModeHTML)
}
