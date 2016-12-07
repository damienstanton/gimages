package main

import "testing"

func TestInit(t *testing.T) {
	if 1 < 0 {
		t.Fatal("kaboom")
	}
	t.Log("✅ fake test has passed")
}

// TODO stash the appropriate API keys on Circle so these tests can be run
// TestPing connects to the custom search API and passes if a 200 OK comes back
// func TestPing(t *testing.T) {
// 	res, err := http.Get(endpoint + "foo" + apikey)
// 	if err != nil {
// 		t.Errorf("problem connecting to api: %v", err)
// 	}

// 	if res.StatusCode != http.StatusOK {
// 		t.Fatal("received non-200 response from API endpoint!")
// 	}
// }

// func TestImageAPI(t *testing.T) {
// 	if err := NewImageSearch("apple"); err != nil {
// 		t.Fatalf("error with retrieving images: %v", err)
// 	}
// }
