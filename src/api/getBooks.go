package api

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/mang0kitty/website/src/helpers"
	"github.com/mang0kitty/website/src/models"
)

type XMLStruct struct {
	ISBN          string `xml:"isbn13"`
	Description   string `xml:"description"`
	AverageRating string `xml:"average_rating"`
}

func FetchXMl(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return []byte{}, fmt.Errorf("GET error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return []byte{}, fmt.Errorf("Status error: %v", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, fmt.Errorf("Read body: %v", err)
	}

	return body, nil
}

func SearchBooks() {

	r, err := os.Open("src/store/books.json")
	helpers.CheckError(err)
	fmt.Println("Successfully Opened books.json")

	defer r.Close()

	byteValue, _ := ioutil.ReadAll(r)

	var books models.Books
	json.Unmarshal(byteValue, &books.Books)

	for _, book := range books.Books {
		GoodReads(book.ISBN)
	}

}

func GoodReads(query string) {
	str := "https://www.goodreads.com/book/isbn/" + query
	u, _ := url.Parse(str)
	q, _ := url.ParseQuery(u.RawQuery)
	q.Add("key", "")
	u.RawQuery = q.Encode()

	if xmlBytes, err := FetchXMl("https://www.goodreads.com/book/isbn/9781524763138?key=NKF8g4vvWIcFVAZtYOsDIA"); err != nil {
		log.Printf("Failed to get XML: %v", err)
	} else {
		var result XMLStruct
		xml.Unmarshal(xmlBytes, &result)
		fmt.Printf("+v\n", result.ISBN)
	}

}
