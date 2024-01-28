package books

import (
	"encoding/json"
	"net/http"

	"github.com/PaoloProdossimoLopes/go-library/database"
	"github.com/PaoloProdossimoLopes/go-library/logger"
	"github.com/PaoloProdossimoLopes/go-library/server"
)

func CreateNewBookHandler(w http.ResponseWriter, r *http.Request) {

	var createNewBookRequest BookRequest
	decodeError := json.NewDecoder(r.Body).Decode(&createNewBookRequest)
	if decodeError != nil {
		logger.Error("Problem to decode create new book request")
		server.SendErrorResponse(w, server.ResponseError{
			Error:      "Internal server error",
			Reason:     decodeError.Error(),
			StatusCode: http.StatusInternalServerError,
		})
		return
	}

	createdBookResponse, createdBookResponseError := database.BooksRepository.CreateNewBook(database.Book{
		Title:  createNewBookRequest.Title,
		Author: createNewBookRequest.Author,
	})
	if createdBookResponseError != nil {
		logger.Error("Problem to create new book")
		server.SendErrorResponse(w, server.ResponseError{
			Error:      "Internal server error",
			Reason:     createdBookResponseError.Error(),
			StatusCode: http.StatusInternalServerError,
		})
		return
	}

	createdBookResponseBytes, booksResponseMarshalError := json.Marshal(
		bookToBookResponseRoot(createdBookResponse))
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
