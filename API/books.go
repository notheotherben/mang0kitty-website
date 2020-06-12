package api

type Book struct {
	Title       string   `json:"title"`
	Name        string   `json:"name"`
	Authour     string   `json:"author"`
	DateRead    string   `json:"dateRead"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
}
