# Install

manual install and launch the server :

    curl -L https://ollama.ai/download/ollama-linux-amd64 -o /usr/local/bin/ollama
    chmod +x /usr/local/bin/ollama
    ollama serve

then in another terminal, run the model :

    ollama run codellama

# Déroulé du test

>>> How to write a program in Go that handles an http POST request with multiple parameters ?

To handle an HTTP POST request with multiple parameters in Go, you can use the `net/http` package. Here is an example of
how you can do this:
```
package main

import (
        "fmt"
        "log"
        "net/http"
)

func main() {
        http.HandleFunc("/post", func(w http.ResponseWriter, r *http.Request) {
                // Check if the request method is POST
                if r.Method != "POST" {
                        log.Println("Invalid request method")
                        return
                }

                // Get the values from the POST request
                values := r.FormValue()

                // Print out the values
                for key, value := range values {
                        fmt.Printf("%s=%s\n", key, value)
                }

                // Write a response back to the client
                w.WriteHeader(http.StatusOK)
        })

        log.Fatal(http.ListenAndServe(":8080", nil))
}
```
In this example, we are using the `net/http` package to handle an HTTP POST request at the endpoint `/post`. When a POST
request is made to this endpoint, we check if the request method is "POST" and then get the values from the request
body. We print out the values using a for loop, and then write a response back to the client with a 200 status code.

You can test this program by running it and making a POST request to `http://localhost:8080/post` with the following
JSON data in the request body:
```
{
        "name": "John Doe",
        "age": 30,
        "email": "johndoe@example.com"
}
```
You should see the values printed out in your console, and a 200 status code returned to the client.

>>> how to use curl to make a post request with JSON data ?

To use `curl` to make a POST request with JSON data, you can use the `-X` or `--request` option followed by the type of
HTTP request (in this case, "POST") and the `-d` or `--data` option followed by the JSON data. Here's an example:
```
curl -X POST -H "Content-Type: application/json" -d '{"name": "John Doe", "age": 30}' https://example.com/api/users
```
In this example, we are making a POST request to the URL `https://example.com/api/users` with the JSON data `{"name":
"John Doe", "age": 30}`. The `-X` or `--request` option specifies that the request should be a POST, and the `-H` or
`--header` option sets the `Content-Type` header to `application/json`, which is required for sending JSON data in a
POST request. The `-d` or `--data` option specifies the JSON data to send with the request.

You can also use the `--data-binary` option to send raw binary data, such as image files, in a POST request. For
example:
```
curl -X POST -H "Content-Type: application/octet-stream" --data-binary @image.jpg https://example.com/api/images
```
In this example, we are sending an image file called `image.jpg` as raw binary data in a POST request to the URL
`https://example.com/api/images`. The `-X` or `--request` option specifies that the request should be a POST, and the
`-H` or `--header` option sets the `Content-Type` header to `application/octet-stream`, which is required for sending
binary data in a POST request. The `--data-binary` option specifies the file to send as raw binary data with the
request.

You can also use the `-F` or `--form` option to send form data, such as key-value pairs, in a POST request. For example:
```
curl -X POST -H "Content-Type: application/x-www-form-urlencoded" -d 'name=John%20Doe&age=30'
https://example.com/api/users
```
In this example, we are sending form data with the key-value pairs `name=John Doe` and `age=30` in a POST request to the
URL `https://example.com/api/users`. The `-X` or `--request` option specifies that the request should be a POST, and the
`-H` or `--header` option sets the `Content-Type` header to `application/x-www-form-urlencoded`, which is required for
sending form data in a POST request. The `-d` or `--data` option specifies the form data to send with the request.

These are just a few examples of how you can use `curl` to make POST requests with different types of data. You can
refer to the `curl` manual for more information on all the options and features that it provides.

>>> Send a message (/? for help)

