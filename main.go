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
		Addr:    "0.0.0.0:8000",

		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Starting server on 0.0.0.0:8000")
	log.Fatal(srv.ListenAndServe())

}
