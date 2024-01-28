package database

import (
	"time"

	"github.com/google/uuid"
)

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

func (r *InMemoryBookRepository) UpdateBook(bk Book) (*Book, error) {
	for index, book := range r.books {
		if book.ID == bk.ID {
			if bk.Title != "" {
				book.Title = bk.Title
			}

			if bk.Author != "" {
				book.Author = bk.Author
			}

			book.UpdatedAt = time.Now().String()

			r.books[index] = book
			return &book, nil
		}
	}

	return nil, nil
}

func (r *InMemoryBookRepository) DeleteBook(id string) (*Book, error) {
	for index, book := range r.books {
		if book.ID == id {
			r.books = append(r.books[:index], r.books[index+1:]...)
			return &book, nil
		}
	}

	return nil, nil
}
