package wellcome

import (
	"encoding/json"
	"net/http"

	"github.com/PaoloProdossimoLopes/go-library/logger"
	"github.com/PaoloProdossimoLopes/go-library/server"
)

type wellcome struct {
	Message string `json:"message"`
}

func GetWellcomeHandler(w http.ResponseWriter, r *http.Request) {
	wellcome := wellcome{"Wellcome to my API"}
	wellcomeData, wellcomeMarshalError := json.Marshal(wellcome)
	if wellcomeMarshalError != nil {
		logger.Error("Problem to marshal wellcome message struct")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(server.ResponseError{
			Error:      "Internal server error",
			Reason:     wellcomeMarshalError.Error(),
			StatusCode: http.StatusInternalServerError,
		}.Marshal())
		return
	}

	w.Write(wellcomeData)
}
