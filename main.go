package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/AIRCentre/webhook-spaceway-lora/external/mysqldriver"
	"github.com/AIRCentre/webhook-spaceway-lora/internal/events_repository"
	"github.com/gorilla/mux"
)

func main() {

	mysqlDriver, err := mysqldriver.New(
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DB_NAME"),
	)
	if err != nil {
		panic(err.Error())
	}

	repo := events_repository.NewMysqlRepo(mysqlDriver)
	uplinkHandlerFunc := buildUplinkHandlerFunc(repo)

	router := mux.NewRouter()
	router.HandleFunc("/health", healthckeckHandlerFunc).Methods("GET")
	router.HandleFunc("/uplink", uplinkHandlerFunc).Methods("POST")

	err = http.ListenAndServe(":3000", router)
	if err != nil {
		panic(err.Error())
	}

}

func healthckeckHandlerFunc(w http.ResponseWriter, r *http.Request) {
	payload := map[string]any{"status": "healthy"}
	json.NewEncoder(w).Encode(payload)
}

func buildUplinkHandlerFunc(repo events_repository.I) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var payload events_repository.SwarmPayload
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
