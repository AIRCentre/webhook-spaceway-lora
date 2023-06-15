package handlers

import (
	"encoding/json"
	"net/http"
)

func NewHealthckeckHandlerFunc() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		payload := map[string]any{"status": "healthy"}
		json.NewEncoder(w).Encode(payload)
	}
}
