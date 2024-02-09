# Install

Utilisation de l'interface web.

# Déroulé du test

User

How to write a program in Go that handles an http POST request with multiple parameters ?

ChatGPT

To write a program in Go that handles an HTTP POST request with multiple parameters, you can use the net/http package along with the http.HandleFunc function to define your HTTP handler. Here's a simple example:

go

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

In this example:

    We define a main function where we set up an HTTP handler using http.HandleFunc.
    The handlePost function is defined to handle POST requests. Inside this function:
        We check if the request method is POST.
        Parse the form data using r.ParseForm().
        Get the values for the parameters using r.Form.Get("param_name").
        Perform any necessary processing on the parameters.
        Respond back to the client with the received parameters.

Make sure to start this program, and it will listen on port 8080. You can then send a POST request with parameters param1 and param2, and the server will handle it accordingly.

User

How to send a POST request with parameters param1 and param2 with curl ?

ChatGPT

You can use curl to send a POST request with parameters param1 and param2 like this:

bash

curl -X POST -d "param1=value1&param2=value2" http://localhost:8080/post

In this command:

    -X POST specifies that the request method is POST.
    -d "param1=value1&param2=value2" specifies the parameters to send in the request body. Adjust the values of param1 and param2 as needed.
    http://localhost:8080/post is the URL to which the request is sent. Adjust the hostname and port if your server is running on a different address or port.
