package extract

import (
	"fmt"
	"tracker-system/pkg/parse"

	"github.com/google/uuid"
	"gopkg.in/telebot.v3"
)

func UUIDFromTeleCtx(c telebot.Context, key string) (uuid.UUID, error) {
	id, err := parse.UUIDFromAny(c.Get(key))
	if err != nil {
		return uuid.Nil, fmt.Errorf("telebot ctx key %q: %w", key, err)
	}
	return id, nil
}
