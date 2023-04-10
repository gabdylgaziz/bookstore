package book

type Book struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Description string `json:"description"`
	Price int `json:"price"`
}
