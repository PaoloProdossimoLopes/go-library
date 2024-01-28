package books

import "github.com/PaoloProdossimoLopes/go-library/database"

func bookToResponse(book database.Book) BookResponse {
	return BookResponse{
		ID:        book.ID,
		Title:     book.Title,
		Author:    book.Author,
		CreatedAt: book.CreatedAt,
		UpdatedAt: book.UpdatedAt,
	}
}

func booksToBooksResponse(books []database.Book) []BookResponse {
	var booksResponse = []BookResponse{}

	for _, book := range books {
		booksResponse = append(booksResponse, bookToResponse(book))
	}

	return booksResponse
}

func booksToBooksResponseRoot(books []database.Book) BooksResponseRoot {
	return BooksResponseRoot{
		Books: booksToBooksResponse(books),
	}
}

func bookToBookResponseRoot(book database.Book) BookResponseRoot {
	return BookResponseRoot{
		Book: bookToResponse(book),
	}
}
