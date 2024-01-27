package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/PaoloProdossimoLopes/go-library/enviroment"
	"github.com/PaoloProdossimoLopes/go-library/logger"
	"github.com/PaoloProdossimoLopes/go-library/wellcome"
)

func main() {
	prepareEnviromentVariables()

	const API = "/api/v1"
	http.HandleFunc(API+"/", Get(wellcome.GetWellcomeHandler))
	http.ListenAndServe(enviroment.Enviroment.GetPort(), nil)
}

func prepareEnviromentVariables() {
	if initEnvError := enviroment.Init(); initEnvError != nil {
		logger.Fatal(initEnvError.Error())
		panic(initEnvError)
	}
}

func Get(handler func(hw http.ResponseWriter, hr *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			logger.Error(fmt.Sprintf(
				"Invalid HTTP method at route '%v' with method '%v'.",
				r.RequestURI, r.Method))
			w.WriteHeader(http.StatusMethodNotAllowed)

			responseErrorData, responseErrorMarshalError := json.Marshal(wellcome.ResponseError{
				Error:      "Method HTTP not allowed",
				Reason:     fmt.Sprintf("Method HTTP '%s' not allowed", r.Method),
				StatusCode: http.StatusMethodNotAllowed,
			})
			if responseErrorMarshalError != nil {
				logger.Error("Problem to marshal `ErrorResponse` struct")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			w.Write([]byte(responseErrorData))
			return
		}

		handler(w, r)
	}
}
