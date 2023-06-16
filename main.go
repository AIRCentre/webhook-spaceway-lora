package main

import (
	"net/http"

	"github.com/AIRCentre/webhook-spaceway-lora/external/restengine"
)

func main() {
	router := restengine.BuildRouter()
	err := http.ListenAndServe(":3000", router)
	if err != nil {
		panic(err.Error())
	}
}
