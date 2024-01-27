package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/PaoloProdossimoLopes/go-library/enviroment"
)

func main() {
	if initEnvError := enviroment.Init(); initEnvError != nil {
		log.Println(initEnvError.Error())
		panic(initEnvError)
	}

	const api = "/api/v1"
	http.HandleFunc(api+"/", getWellcome)
	http.ListenAndServe(enviroment.Enviroment.GetPort(), nil)
}

type Wellcome struct {
	Message string `json:"message"`
}

func getWellcome(w http.ResponseWriter, r *http.Request) {
	wellcome := Wellcome{"Wellcome to my API"}
	wellcomeData, wellcomeMarshalError := json.Marshal(wellcome)
	if wellcomeMarshalError != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(wellcomeData)
}
