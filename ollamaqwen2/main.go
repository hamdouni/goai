package main

import (
	"fmt"
	"log"
	"net/http"

	// This line imports the url.Values type for handling form data.
	_ "net/url"
)

// This function will handle the POST request and process its parameters
func HandlePost(w http.ResponseWriter, r *http.Request) {
	// Check if this request is indeed a POST request
	if r.Method != http.MethodPost {
		http.Error(w, "This endpoint only accepts HTTP POST requests", http.StatusMethodNotAllowed)
		return
	}

	params := make([]string, 0)

	// Get form data from the request.
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		http.Error(w, "An error occurred during parsing", http.StatusInternalServerError)
		return
	}

	for k, vs := range r.Form {
		fmt.Fprintf(w, "%s = %v\n", k, vs)
		params = append(params, k+"="+vs[0])
	}
	// We can process the `params` array here if necessary.
}

func main() {
	http.HandleFunc("/", HandlePost) // Set our handler function to be called when `/` is requested.

	log.Println("Server starting on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
