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
