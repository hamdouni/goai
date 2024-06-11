package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "ParseForm() error: %v", err)
		return
	}

	name := r.PostFormValue("name")
	email := r.PostFormValue("email")

	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Email = %s\n", email)
}

func main() {
	http.HandleFunc("/submit", handler) // If you want to handle multiple parameters, just add more routes and handlers here

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
