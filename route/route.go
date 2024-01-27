package route

import (
	"github.com/PaoloProdossimoLopes/go-library/books"
	"github.com/PaoloProdossimoLopes/go-library/server"
	"github.com/PaoloProdossimoLopes/go-library/wellcome"
)

func Init() {
	server.Get("/", wellcome.GetWellcomeHandler)

	server.Get("/books", books.GetAllBooksHandler)
	server.Post("/books", books.CreateNewBookHandler)
}
