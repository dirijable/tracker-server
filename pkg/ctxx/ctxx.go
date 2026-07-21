package ctxx

import (
	"context"

	"gopkg.in/telebot.v3"
)

var key = "ctx"

func Set(c telebot.Context, ctx context.Context) {
	c.Set(key, ctx)
}

func From(c telebot.Context) context.Context {
	if ctx, ok := c.Get(key).(context.Context); ok {
		return ctx
	}
	return context.Background()
}
