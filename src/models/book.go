package models

type Book struct {
	ID          int      `json:"id"`
	Title       string   `json:"title"`
	Authour     []string `json:"author"`
	DateRead    string   `json:"date_read"`
	Description string   `json:"description"`
	Photo       string   `json:"photo"`
	ISBN        string   `json:"isbn"`
	Tags        []string `json:"tags"`
}
