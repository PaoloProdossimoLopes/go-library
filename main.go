package main

import (
	"encoding/json"
	"net/http"

	"github.com/PaoloProdossimoLopes/go-library/enviroment"
	"github.com/PaoloProdossimoLopes/go-library/logger"
)

func main() {
	prepareEnviromentVariables()

	const API = "/api/v1"
	http.HandleFunc(API+"/", getWellcomeHandler)
	http.ListenAndServe(enviroment.Enviroment.GetPort(), nil)
}

func prepareEnviromentVariables() {
	if initEnvError := enviroment.Init(); initEnvError != nil {
		logger.Fatal(initEnvError.Error())
		panic(initEnvError)
	}
}

type wellcome struct {
	Message string `json:"message"`
}

func getWellcomeHandler(w http.ResponseWriter, r *http.Request) {
	wellcome := wellcome{"Wellcome to my API"}
	wellcomeData, wellcomeMarshalError := json.Marshal(wellcome)
	if wellcomeMarshalError != nil {
		logger.Error("Problem to marshal wellcome message struct")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(wellcomeData)
}
