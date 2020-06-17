package api

// type XMLStruct struct {
// 	ISBN          string `xml:"isbn13"`
// 	Description   string `xml:"description"`
// 	AverageRating string `xml:"average_rating"`
// }

// func FetchXML(url string) ([]byte, error) {
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		return []byte{}, fmt.Errorf("GET error: %v", err)
// 	}
// 	defer resp.Body.Close()

// 	if resp.StatusCode != http.StatusOK {
// 		return []byte{}, fmt.Errorf("Status error: %v", resp.StatusCode)
// 	}

// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		return []byte{}, fmt.Errorf("Read body: %v", err)
// 	}

// 	return body, nil
// }
