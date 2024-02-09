package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/post", handlePost)
	fmt.Println("Server is listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	// Check if the HTTP method is POST
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse form data
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusInternalServerError)
		return
	}

	// Get values for multiple parameters
	param1 := r.Form.Get("param1")
	param2 := r.Form.Get("param2")

	// You can access more parameters similarly

	// Do something with the parameters
	fmt.Fprintf(w, "Received parameters:\n")
	fmt.Fprintf(w, "param1: %s\n", param1)
	fmt.Fprintf(w, "param2: %s\n", param2)

	// You can process the parameters further as needed
}
