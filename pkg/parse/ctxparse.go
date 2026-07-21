package parse

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
)

func UUIDFromAny(value any) (uuid.UUID, error) {
	if value == nil {
		return uuid.Nil, errors.New("value is nil")
	}
	switch v := value.(type) {
	case uuid.UUID:
		return v, nil
	case string:
		uuidParsed, err := uuid.Parse(v)
		if err != nil {
			return uuid.Nil, fmt.Errorf("parse string to uuid: %w", err)
		}
		return uuidParsed, nil
	default:
		return uuid.Nil, fmt.Errorf("unexpected type: %T", v)
	}
}
