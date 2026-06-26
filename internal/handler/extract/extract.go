package extract

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

func ExtractUUID(ctx context.Context, key string) (uuid.UUID, error) {
	value := ctx.Value(key)
	if value == nil {
		return uuid.Nil, fmt.Errorf("key %q not found in context", key)
	}
	uuidStr, ok := value.(string)
	if !ok {
		return uuid.Nil, fmt.Errorf("value by key %q is not a string", key)
	}
	uuidParsed, err := uuid.Parse(uuidStr)
	if err != nil {
		return uuid.Nil, fmt.Errorf("parse uuid: %w", err)
	}
	return uuidParsed, nil
}
