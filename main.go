package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	http.HandleFunc("/", getWellcome)
	http.ListenAndServe(":8080", nil)
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
