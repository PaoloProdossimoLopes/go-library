package database

var BooksRepository BookRepository

func Init() {
	BooksRepository = &InMemoryBookRepository{}
}
