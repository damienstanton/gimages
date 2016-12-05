package main

import (
	"net/http"
	"testing"
)

// TestPing connects to the custom search API and passes if a 200 OK comes back
func TestPing(t *testing.T) {
	res, err := http.Get(endpoint + "foo" + apikey)
	if err != nil {
		t.Errorf("problem connecting to api: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Fatal("received non-200 response from API endpoint!")
	}
}

func TestImageAPI(t *testing.T) {
	if _, err := NewImageSearch("apple"); err != nil {
		t.Fatalf("error with retrieving images: %v", err)
	}
}
