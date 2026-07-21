package extract

import (
	"context"
	"fmt"
	"tracker-system/pkg/parse"

	"github.com/google/uuid"
)

func UUIDFromCtx(ctx context.Context, key any) (uuid.UUID, error) {
	id, err := parse.UUIDFromAny(ctx.Value(key))
	if err != nil {
		return uuid.Nil, fmt.Errorf("context key %v: %w", key, err)
	}
	return id, nil
}
