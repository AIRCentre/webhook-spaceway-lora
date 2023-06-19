package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/AIRCentre/webhook-spaceway-lora/internal/eventrepo"
)

func NewUplinkHandlerFunc(repo eventrepo.I) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read request body", http.StatusInternalServerError)
			return
		}
		fmt.Printf("Request Body: %s", body)
		reader := io.NopCloser(bytes.NewBuffer(body))
		var payload eventrepo.EventPayload
		err = json.NewDecoder(reader).Decode(&payload)
		if err != nil {
			fmt.Println(err.Error())
			http.Error(w, "Failed to read request body", http.StatusInternalServerError)
			return
		}
		err = repo.Insert(payload)
		if err != nil {
			fmt.Println(err.Error())
			http.Error(w, "Failed to handle swarm payload", http.StatusInternalServerError)
			return
		}
	}
}
