package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/AIRCentre/webhook-spaceway-lora/external/mysqldriver"
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

	fmt.Println(mysqlDriver) // remove this after driver is in use

	router := mux.NewRouter()
	router.HandleFunc("/health", healthckeckHandlerFunc).Methods("GET")

	err = http.ListenAndServe(":3000", router)
	if err != nil {
		panic(err.Error())
	}

}

func healthckeckHandlerFunc(w http.ResponseWriter, r *http.Request) {
	payload := map[string]any{"status": "healthy"}
	json.NewEncoder(w).Encode(payload)
}
