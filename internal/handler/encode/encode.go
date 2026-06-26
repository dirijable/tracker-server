package encode

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func SendJSONResponse(w http.ResponseWriter, statusCode int, body any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(body); err != nil {
		return fmt.Errorf("send json response: %w", err)
	}
	return nil
}
