package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleHomeRoute(t *testing.T) {
	// set up a new request
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// set up a response recorder
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HandleHomeRoute)

	// call the handler
	handler.ServeHTTP(rr, req)

	// tests
	// test status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code; Wanted %v Got %v", status, http.StatusOK)
	}

	// check content type
	if ctype := rr.Header().Get("Content-Type"); ctype != "text/plain" {
		t.Errorf("Handler returned wrong content Type: Wanted text/plain, Got %v", ctype)
	}

	// parse json response
	var paths map[string]Path

	err = json.NewDecoder(rr.Body).Decode(&paths)
	if err != nil {
		t.Errorf("Failed to parse JSON; %v", err)
	}

	// test expected paths
	expectedPaths := []string{"Home", "Hello"}

	for _, pathName := range expectedPaths {
		if _, exists := paths[pathName]; !exists {
			t.Errorf("Expected path %v not found", pathName)
		}
	}
}

func TestHandleHello(t *testing.T) {
	// Create a request to pass to our handler
	req, err := http.NewRequest("GET", "/hello", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HandleHello)

	// Call the handler
	handler.ServeHTTP(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body
	expected := "World"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
