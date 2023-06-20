package handlers

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
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
		log.Printf("payload recieved!\n%s", body)
		reader := io.NopCloser(bytes.NewBuffer(body))

		eventMap := map[string]any{}
		err = json.NewDecoder(reader).Decode(&eventMap)
		if err != nil {
			fmt.Println(err.Error())
			http.Error(w, "Failed to read request body", http.StatusInternalServerError)
			return
		}
		eventData, err := decodeBase64(eventMap["data"].(string))
		if err != nil {
			fmt.Println(err.Error())
			http.Error(w, "Failed to decode device data", http.StatusInternalServerError)
			return
		}
		var deviceData eventrepo.EventPayload
		json.Unmarshal([]byte(eventData), &deviceData)

		err = repo.Insert(fmt.Sprint(eventMap["deviceId"]), deviceData)
		if err != nil {
			fmt.Println(err.Error())
			http.Error(w, "Failed to handle swarm payload", http.StatusInternalServerError)
			return
		}
	}
}

func decodeBase64(encodedString string) (string, error) {
	decodedBytes, err := base64.StdEncoding.DecodeString(encodedString)
	if err != nil {
		return "", err
	}
	return string(decodedBytes), nil
}
