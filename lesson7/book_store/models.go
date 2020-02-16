package book_store

type BookStoreClass struct {
	Name       string  `json:"name"`
	Books      []*Book `json:"books"`
	LastBookID int
}

type Book struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Author      string `json:"author"`
}

type Config struct {
	PathToBookStore string `json:"pathToBookStore"`
	Port            string `json:"port"`
}
