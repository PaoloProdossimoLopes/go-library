package wellcome

import (
	"encoding/json"
	"net/http"

	"github.com/PaoloProdossimoLopes/go-library/logger"
)

type wellcome struct {
	Message string `json:"message"`
}

type ResponseError struct {
	Error      string `json:"error"`
	Reason     string `json:"reason"`
	StatusCode int    `json:"status_code"`
}

func GetWellcomeHandler(w http.ResponseWriter, r *http.Request) {
	wellcome := wellcome{"Wellcome to my API"}
	wellcomeData, wellcomeMarshalError := json.Marshal(wellcome)
	if wellcomeMarshalError != nil {
		logger.Error("Problem to marshal wellcome message struct")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(wellcomeData)
}
