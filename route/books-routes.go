package route

import (
	"github.com/PaoloProdossimoLopes/go-library/books"
	"github.com/PaoloProdossimoLopes/go-library/server"
)

func configureBooksRoutes() {
	const booksPath = "/books"
	server.Get(booksPath, books.GetAllBooksHandler)
	server.Patch(booksPath, books.UpdateBookHandler)
	server.Post(booksPath, books.CreateNewBookHandler)
}
