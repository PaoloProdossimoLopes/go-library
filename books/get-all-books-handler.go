package books

import (
	"encoding/json"
	"net/http"

	"github.com/PaoloProdossimoLopes/go-library/logger"
	"github.com/PaoloProdossimoLopes/go-library/server"
)

type GetAllBooksResponse struct {
	Books []BookResponse `json:"books"`
}

func GetAllBooksHandler(w http.ResponseWriter, r *http.Request) {
	booksResponseModel := GetAllBooksResponse{
		Books: []BookResponse{},
	}

	booksResponseBytes, booksResponseMarshalError := json.Marshal(booksResponseModel)
	if booksResponseMarshalError != nil {
		logger.Error("Problem to marshal books response model")
		server.SendErrorResponse(w, server.ResponseError{
			Error:      "Internal server error",
			Reason:     booksResponseMarshalError.Error(),
			StatusCode: http.StatusInternalServerError,
		})
		return
	}

	w.Write(booksResponseBytes)
}
