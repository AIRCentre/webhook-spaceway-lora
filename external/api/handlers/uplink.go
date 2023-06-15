package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/AIRCentre/webhook-spaceway-lora/internal/eventrepo"
)

func NewUplinkHandlerFunc(repo eventrepo.I) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var payload eventrepo.SwarmPayload
		err := json.NewDecoder(r.Body).Decode(&payload)
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
