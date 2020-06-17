package api

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/mang0kitty/website/src/helpers"
	"github.com/mang0kitty/website/src/models"
)

type XMLBook struct {
	ISBN          string  `xml:"book>isbn13"`
	Description   string  `xml:"book>description"`
	AverageRating float32 `xml:"book>average_rating"`
}

func FetchXML(url string) (io.ReadCloser, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("GET error: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Status error: %v", resp.StatusCode)
	}

	return resp.Body, nil

	// body, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	return []byte{}, fmt.Errorf("Read body: %v", err)
	// }

	// return body, nil
}

func SearchBooks() {

	r, err := os.OpenFile("src/store/books.json", os.O_RDWR, os.ModePerm)
	helpers.CheckError(err)
	fmt.Println("Successfully Opened books.json")

	defer r.Close()

	var books []*models.Book
	// TODO: handle errors here
	helpers.CheckError(json.NewDecoder(r).Decode(&books))

	for _, book := range books {
		grb, err := GoodReads(book.ISBN)
		if err != nil {
			fmt.Println("Failed to get info for ISBN", book.ISBN)
			continue
		}

		book.Description = grb.Description
		book.Rating = grb.AverageRating
	}

	// for _, book := range books {
	// 	fmt.Printf("ISBN %s (%f): %s\n", book.ISBN, book.Rating, book.Description)
	// }

	// TODO: handle errors on all of these
	_, err = r.Seek(0, os.SEEK_SET)
	helpers.CheckError(err)
	helpers.CheckError(r.Truncate(0))
	helpers.CheckError(json.NewEncoder(r).Encode(&books))
}

func GoodReads(query string) (*XMLBook, error) {
	str := "https://www.goodreads.com/book/isbn/" + query
	u, _ := url.Parse(str)
	q, _ := url.ParseQuery(u.RawQuery)
	q.Add("key", os.Getenv("GOODREADS_API_KEY"))
	u.RawQuery = q.Encode()

	if xmlStream, err := FetchXML(u.String()); err != nil {
		return nil, err
	} else {
		defer xmlStream.Close()
		var result XMLBook
		if err := xml.NewDecoder(xmlStream).Decode(&result); err != nil {
			return nil, err
		} else {
			return &result, nil
		}
	}
}
