package main

import (
	"net/http"

	"github.com/AIRCentre/webhook-spaceway-lora/external/api/engine"
)

func main() {

	router := engine.Init()

	err := http.ListenAndServe(":3000", router)
	if err != nil {
		panic(err.Error())
	}

}
