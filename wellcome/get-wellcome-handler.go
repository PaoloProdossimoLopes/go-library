package wellcome

import (
	"encoding/json"
	"net/http"

	"github.com/PaoloProdossimoLopes/go-library/logger"
	"github.com/PaoloProdossimoLopes/go-library/server"
)

func GetWellcomeHandler(w http.ResponseWriter, r *http.Request) {
	wellcome := struct {
		Message string `json:"message"`
	}{"Wellcome to my API"}
	wellcomeData, wellcomeMarshalError := json.Marshal(wellcome)
	if wellcomeMarshalError != nil {
		logger.Error("Problem to marshal wellcome message struct")
		server.SendErrorResponse(w, server.ResponseError{
			Error:      "Internal server error",
			Reason:     wellcomeMarshalError.Error(),
			StatusCode: http.StatusInternalServerError,
		})
		return
	}

	w.Write(wellcomeData)
}
