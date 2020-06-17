package main

import (
	"github.com/mang0kitty/website/src/api"
)

func main() {

	// r := handlers.Handle()
	// srv := &http.Server{
	// 	Handler: r,
	// 	Addr:    "127.0.0.1:8000",

	// 	WriteTimeout: 15 * time.Second,
	// 	ReadTimeout:  15 * time.Second,
	// }
	api.SearchBooks()
	// log.Fatal(srv.ListenAndServe())

}
