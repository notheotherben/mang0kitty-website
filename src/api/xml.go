package api

import (
	"fmt"
	"io"
	"net/http"
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

}
