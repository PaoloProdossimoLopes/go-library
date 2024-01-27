package main

import (
	"encoding/json"
	"fmt"
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

type ResponseError struct {
	Error      string `json:"error"`
	Reason     string `json:"reason"`
	StatusCode int    `json:"status_code"`
}

func getWellcomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		logger.Error(fmt.Sprintf(
			"Invalid HTTP method at route '%v' with method '%v'.",
			r.RequestURI, r.Method))
		notAllowedStatusCode := http.StatusMethodNotAllowed
		w.WriteHeader(notAllowedStatusCode)

		responseErrorData, responseErrorMarshalError := json.Marshal(ResponseError{
			Error:      "Method HTTP not allowed",
			Reason:     fmt.Sprintf("Method HTTP '%s' not allowed", r.Method),
			StatusCode: notAllowedStatusCode,
		})
		if responseErrorMarshalError != nil {
			logger.Error("Problem to marshal `ErrorResponse` struct")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write([]byte(responseErrorData))
		return
	}

	wellcome := wellcome{"Wellcome to my API"}
	wellcomeData, wellcomeMarshalError := json.Marshal(wellcome)
	if wellcomeMarshalError != nil {
		logger.Error("Problem to marshal wellcome message struct")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(wellcomeData)
}
