package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type ParamsStruct struct {
	Param1 string `json:"param1"`
	Param2 int    `json:"param2"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	var params ParamsStruct
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	// Print the received parameters for demonstration purpose
	log.Printf("Received parameters: %v\n", params)

	// Create a simple JSON response
	response := map[string]string{"status": "success"}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Write the JSON response to the client
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func main() {
	http.HandleFunc("/post-endpoint", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
