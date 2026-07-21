package bot

import (
	"context"
	"fmt"
	"tracker-system/pkg/ctxx"

	"gopkg.in/telebot.v3"
)

type UserService interface {
	Save(ctx context.Context, tgChatID int64) error
}

type UserHandler struct {
	svc UserService
}

func NewUserHandler(svc UserService) *UserHandler {
	return &UserHandler{
		svc: svc,
	}
}

func (h *UserHandler) Start(ctx telebot.Context) error {
	tgChatID := ctx.Sender().ID
	if err := h.svc.Save(ctxx.From(ctx), tgChatID); err != nil {
		//TODO кастомная обработка ошибок
		return ctx.Send(fmt.Errorf("register new user: %w", err))
	}
	return ctx.Send("Регистрация прошла успешно, пользователь сохранен")
}
