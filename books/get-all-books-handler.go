package books

import (
	"encoding/json"
	"net/http"

	"github.com/PaoloProdossimoLopes/go-library/database"
	"github.com/PaoloProdossimoLopes/go-library/logger"
	"github.com/PaoloProdossimoLopes/go-library/server"
)

type GetAllBooksResponse struct {
	Books []BookResponse `json:"books"`
}

func GetAllBooksHandler(w http.ResponseWriter, r *http.Request) {
	books, booksResponseError := database.BooksRepository.GetAllBooks()
	if booksResponseError != nil {
		logger.Error("Problem to get all books")
		server.SendErrorResponse(w, server.ResponseError{
			Error:      "Internal server error",
			Reason:     booksResponseError.Error(),
			StatusCode: http.StatusInternalServerError,
		})
		return
	}

	booksResponseBytes, booksResponseMarshalError := json.Marshal(
		booksToBooksResponseRoot(books))
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
