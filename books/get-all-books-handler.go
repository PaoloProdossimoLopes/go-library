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

	booksResponseBytes, booksResponseMarshalError := json.Marshal(GetAllBooksResponse{
		Books: mapBooksToBooksResponse(books),
	})
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

func mapBooksToBooksResponse(books []database.Book) []BookResponse {
	var booksResponse = []BookResponse{}

	for _, book := range books {
		booksResponse = append(booksResponse, BookResponse{
			ID:        book.ID,
			Title:     book.Title,
			Author:    book.Author,
			CreatedAt: book.CreatedAt,
			UpdatedAt: book.UpdatedAt,
		})
	}

	return booksResponse
}
