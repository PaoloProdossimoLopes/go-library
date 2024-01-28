package books

import (
	"encoding/json"
	"net/http"

	"github.com/PaoloProdossimoLopes/go-library/database"
	"github.com/PaoloProdossimoLopes/go-library/logger"
	"github.com/PaoloProdossimoLopes/go-library/server"
)

type DeleteBookResponse struct {
	Book BookResponse `json:"book"`
}

func DeleteBookHandler(w http.ResponseWriter, r *http.Request) {
	bookId := r.URL.Query().Get("id")
	if bookId == "" {
		logger.Error("Problem to get book id")
		server.SendErrorResponse(w, server.ResponseError{
			Error:      "Bad Request",
			Reason:     "book id is missing",
			StatusCode: http.StatusBadRequest,
		})
		return
	}

	updatedBookResponse, updatedBookResponseError := database.BooksRepository.DeleteBook(bookId)
	if updatedBookResponseError != nil {
		logger.Error("Problem to update book")
		server.SendErrorResponse(w, server.ResponseError{
			Error:      "Internal server error",
			Reason:     updatedBookResponseError.Error(),
			StatusCode: http.StatusInternalServerError,
		})
		return
	}

	updatedBookResponseBytes, booksResponseMarshalError := json.Marshal(
		bookToBookResponseRoot(*updatedBookResponse))
	if booksResponseMarshalError != nil {
		logger.Error("Problem to marshal books response model")
		server.SendErrorResponse(w, server.ResponseError{
			Error:      "Internal server error",
			Reason:     booksResponseMarshalError.Error(),
			StatusCode: http.StatusInternalServerError,
		})
		return
	}

	w.Write(updatedBookResponseBytes)
}
