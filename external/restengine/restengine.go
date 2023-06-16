package restengine

import (
	"os"

	"github.com/AIRCentre/webhook-spaceway-lora/external/mysqldriver"
	"github.com/AIRCentre/webhook-spaceway-lora/external/restengine/handlers"
	"github.com/AIRCentre/webhook-spaceway-lora/internal/eventrepo"
	"github.com/gorilla/mux"
)

func BuildRouter() *mux.Router {
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

	repo := eventrepo.NewMysqlRepo(mysqlDriver)
	healthckeckHandlerFunc := handlers.NewHealthckeckHandlerFunc()
	uplinkHandlerFunc := handlers.NewUplinkHandlerFunc(repo)

	router := mux.NewRouter()
	router.HandleFunc("/health", healthckeckHandlerFunc).Methods("GET")
	router.HandleFunc("/uplink", uplinkHandlerFunc).Methods("POST")

	return router

}
