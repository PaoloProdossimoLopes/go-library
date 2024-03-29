package books

type BookResponseRoot struct {
	Book BookResponse `json:"book"`
}

type BooksResponseRoot struct {
	Books []BookResponse `json:"books"`
}

type BookResponse struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
