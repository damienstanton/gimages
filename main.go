package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/jmoiron/jsonq"
)

const endpoint = "https://www.googleapis.com/customsearch/v1?q="

var (
	apikey = os.Getenv("IMGKEY")
	csekey = os.Getenv("CSEKEY")
)

// Result holds response data
type Result struct {
	Items []struct {
		Pagemap struct {
			Imageobject []struct {
				URL string
			}
		}
	}
}

func main() {
	fmt.Println("TODO")
	err := NewImageSearch("apple")
	if err != nil {
		log.Fatal("problem")
	}
}

// NewImageSearch dials Google via the custom search REST API
func NewImageSearch(q string) error {
	k := keyCheck()
	if !k {
		log.Fatal("Exiting due to missing config vars")
	}

	req, err := http.NewRequest("GET", endpoint+q, nil)
	if err != nil {
		log.Errorf("problem connecting to Google CSE: %v", err)
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Errorf("problem opening connection: %v", err)
	}

	defer res.Body.Close()

	var result Result
	jres := json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		log.Errorf("problem decoding result: %v", err)
	}

	// TODO iterate over each item and pull out image URL
	jq := jsonq.NewQuery(jres)
	jq.String("Items", "Pagemap", "Imageobject", "URL")

	// copied from:
	// https://github.com/thbar/golang-playground/blob/master/download-files.go
	tokens := strings.Split(endpoint+q, "/")
	fileName := tokens[len(tokens)-1]
	file, err := os.Create(fileName)
	if err != nil {
		log.Errorf("problem writing file: %v", err)
	}

	defer file.Close()

	if _, err := io.Copy(file, res.Body); err != nil {
		log.Errorf("problem downloading data to file: %v", err)
	}

	return nil
}

// keyCheck ensures that the appropriate API and CSE key/ids are in place
func keyCheck() bool {
	if apikey == "" {
		log.Fatal("You need to export the IMGKEY environment variable")
		return false
	}
	if csekey == "" {
		log.Fatal("You need to export the CSEKEY environment variable")
		return false
	}
	return true
}
