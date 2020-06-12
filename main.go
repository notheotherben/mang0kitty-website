package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}
func bookHandler(w http.ResponseWriter, r *http.Request) {
	return
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/books", bookHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
