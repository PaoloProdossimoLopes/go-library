package database

import (
	"time"

	"github.com/google/uuid"
)

type BookRepository interface {
	CreateNewBook(book Book) (Book, error)
	GetAllBooks() ([]Book, error)
}

type Book struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

var BooksRepository BookRepository

func Init() {
	BooksRepository = &InMemoryBookRepository{}
}

type InMemoryBookRepository struct {
	books []Book
}

func (r *InMemoryBookRepository) CreateNewBook(book Book) (Book, error) {
	book.ID = uuid.New().String()
	book.CreatedAt = time.Now().String()
	book.UpdatedAt = time.Now().String()

	r.books = append(r.books, book)

	return book, nil
}

func (r *InMemoryBookRepository) GetAllBooks() ([]Book, error) {
	return r.books, nil
}
