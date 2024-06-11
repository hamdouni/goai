package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandlePost(t *testing.T) {
	t.Run("Success with all fields filled", func(t *testing.T) {
		testServer := httptest.NewServer(http.HandlerFunc(HandlePost))
		defer testServer.Close()

		client := &http.Client{}
		req, err := http.NewRequest("POST", testServer.URL, strings.NewReader(`
                        firstParam=hello&secondParam=world
                `))
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		resp, err := client.Do(req)
		if err != nil {
			t.Error(err)
			return
		}

		defer resp.Body.Close()

		body, _ := ioutil.ReadAll(resp.Body)

		if string(body) != `firstParam = hello\nsecondParam = world\n` {
			t.Errorf("Unexpected response body: %s", string(body))
		}
	})

	t.Run("No fields filled in POST request", func(t *testing.T) {
		testServer := httptest.NewServer(http.HandlerFunc(HandlePost))
		defer testServer.Close()

		client := &http.Client{}
		req, err := http.NewRequest("POST", testServer.URL, strings.NewReader(""))
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		resp, err := client.Do(req)
		if err != nil {
			t.Error(err)
		}

		defer resp.Body.Close()

		body, _ := ioutil.ReadAll(resp.Body)

		if resp.StatusCode != http.StatusNotAcceptable {
			t.Errorf("Response status code mismatch: %d", resp.StatusCode)
		}
		if string(body) != `This endpoint only accepts HTTP POST requests` {
			t.Errorf("Unexpected response body: %s", string(body))
		}
	})

	t.Run("GET request", func(t *testing.T) {
		testServer := httptest.NewServer(http.HandlerFunc(HandlePost))
		defer testServer.Close()

		client := &http.Client{}
		req, err := http.NewRequest("GET", testServer.URL, nil)
		if err != nil {
			t.Fatal(err)
		}

		resp, err := client.Do(req)
		if err != nil {
			t.Error(err)
		}

		defer resp.Body.Close()

		body, _ := ioutil.ReadAll(resp.Body)

		if resp.StatusCode != http.StatusMethodNotAllowed {
			t.Errorf("Response status code mismatch: %d", resp.StatusCode)
		}
		if string(body) != `This endpoint only accepts HTTP POST requests` {
			t.Errorf("Unexpected response body: %s", string(body))
		}
	})
}
