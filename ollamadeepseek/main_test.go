package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	type testStruct struct {
		Param1 string `json:"param1"`
		Param2 int    `json:"param2"`
	}

	tcs := []struct {
		name       string
		body       testStruct
		wantStatus int
	}{
		{
			"valid request",
			testStruct{"value1", 2},
			http.StatusOK,
		},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			// Convert the test struct into a JSON string.
			jsonStr, _ := json.Marshal(tc.body)

			// Create a new request with JSON body.
			req, err := http.NewRequest("POST", "/post-endpoint", bytes.NewBufferString(string(jsonStr)))
			if err != nil {
				t.Fatalf("Could not create request: %v\n", err)
			}

			// Create a new ResponseRecorder which satisfies http.ResponseWriter.
			rec := httptest.NewRecorder()
			handler(rec, req)

			// Check the status code is what we expect.
			if status := rec.Code; status != tc.wantStatus {
				t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
			}
		})
	}
}
