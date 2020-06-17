package main

import (
	"log"
	"net/http"
	"time"

	"github.com/mang0kitty/website/src/handlers"
)

func main() {
	// api.SearchBooks()

	r := handlers.Handle()
	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",

		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())

}
