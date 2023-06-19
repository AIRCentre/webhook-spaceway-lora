package main

import (
	"log"
	"net/http"

	"github.com/AIRCentre/webhook-spaceway-lora/external/restengine"
)

func main() {
	router := restengine.BuildRouter()
	log.Println("service running on port 3000")
	err := http.ListenAndServe(":3000", router)
	if err != nil {
		panic(err.Error())
	}

}
