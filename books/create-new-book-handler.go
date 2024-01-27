package books

import (
	"encoding/json"
	"net/http"

	"github.com/PaoloProdossimoLopes/go-library/logger"
	"github.com/PaoloProdossimoLopes/go-library/server"
)

type CreateNewBookResponse struct {
	Book BookResponse `json:"book"`
}

func CreateNewBookHandler(w http.ResponseWriter, r *http.Request) {
	createdBookResponse := CreateNewBookResponse{
		Book: BookResponse{},
	}

	createdBookResponseBytes, booksResponseMarshalError := json.Marshal(createdBookResponse)
	if booksResponseMarshalError != nil {
		logger.Error("Problem to marshal books response model")
		server.SendErrorResponse(w, server.ResponseError{
			Error:      "Internal server error",
			Reason:     booksResponseMarshalError.Error(),
			StatusCode: http.StatusInternalServerError,
		})
		return
	}

	w.Write(createdBookResponseBytes)
}
