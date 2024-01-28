package database

type BookRepository interface {
	CreateNewBook(book Book) (Book, error)
	GetAllBooks() ([]Book, error)
	UpdateBook(book Book) (*Book, error)
	DeleteBook(id string) (*Book, error)
}
