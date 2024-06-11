package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler(t *testing.T) {
	data := "name=John&email=john@example.com"
	req := httptest.NewRequest("POST", "/submit", strings.NewReader(data))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Create a ResponseRecorder to capture handler's output
	rr := httptest.NewRecorder()

	// Execute the request
	handler(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "Name = John\nEmail = john@example.com" + "\n"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
