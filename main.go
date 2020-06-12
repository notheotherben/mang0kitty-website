package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/mang0kitty/website/src/helpers"
)

func BookHandler(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)

	fptr, err := os.Open("src/store/books.json")
	helpers.CheckError(err)
	fmt.Println("Successfully Opened books.json")
	defer fptr.Close()

	byteValue, _ := ioutil.ReadAll(fptr)

	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	json.NewEncoder(w).Encode(result["books"])

}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/books", BookHandler)

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",

		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

}
