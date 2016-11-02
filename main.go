package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/jmoiron/jsonq"
)

const (
	apikey   = os.Getenv("IMGKEY")
	csekey   = os.Getenv("CSEKEY")
	endpoint = "https://www.googleapis.com/customsearch/v1?q="
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

// NewImageSearch dials Google via the custom search REST API
func NewImageSearch(q string) error {
	if keyCheck() == false {
		logrus.Error("Exiting due to missing config vars")
		return err
	}

	req, err := http.NewRequest("GET", endpoint+q, nil)
	if err != nil {
		logrus.Errorf("problem connecting to Google CSE: %v", err)
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		logrus.Errorf("problem opening connection: %v", err)
	}

	defer res.Body.Close()

	var result Result
	jres := json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		logrus.Errorf("problem decoding result: %v", err)
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
		logrus.Errorf("problem writing file: %v", err)
	}

	defer file.Close()

	if err := io.Copy(file, res.Body); err != nil {
		logrus.Errorf("problem downloading data to file: %v", err)
	}

	return nil
}

// keyCheck ensures that the appropriate API and CSE key/ids are in place
func keyCheck() bool {
	if apikey == "" {
		logrus.Fatal("You need to export the IMGKEY environment variable")
		return false
	}
	if csekey == "" {
		logrus.Fatal("You need to export the CSEKEY environment variable")
		return false
	}
	return true
}
