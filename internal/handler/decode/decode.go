package decode

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Decode(r *http.Request, dest any) error {
	if err := json.NewDecoder(r.Body).Decode(dest); err != nil {
		return fmt.Errorf("json decode: %w", err)
	}
	return nil
}
