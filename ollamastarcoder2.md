# Déroulé du test avec le model 7b

- **avis général** :
    - les réponses sont super longues à obtenir (~2mn)
    - les résultats sont plutôt bons, du premier coup on arrive à une réponse correcte
    - sauf pour le test, à un saut de ligne près.


>>> How to write a program in Go that handles an http POST request with multiple parameters ?

In Go, you can handle HTTP POST requests by creating an HTTP server and defining handlers for different URL paths. Here's an example of how you might set up a simple HTTP
server to accept a POST request containing multiple parameters:

```go
package main

import (
    "fmt"
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
```
In this example, the handler function is set to handle POST requests at "/submit". It parses the form data from the request and retrieves the values for "name" and "email"
using `r.PostFormValue()`. Then it sends a response back with these values in plain text format.

To run this server, you would save the code to a file (for example, `main.go`), then compile and run the program:

```bash
$ go run main.go
Starting server at port 8080
```

Now if you make a POST request with form data like so:

```bash
curl -d "name=John&email=john@example.com" -X POST http://localhost:8080/submit
```

The server will respond with:

```text
Name = John
Email = john@example.com
```


>>> Can you provide some tests for this program ?

Testing an HTTP server in Go can be a bit tricky because it involves setting up a
testing server and making actual requests to that server. The standard
`net/http/httptest` package is used to facilitate the creation of such servers.

Here's how you could write some tests for your handler:

```go
package main

import (
        "io/ioutil"
        "net/http/httptest"
        "strings"
        "testing"
)

func TestHandler(t *testing.T) {
        data := "name=John&email=john@example.com"

        // Create a new request
        req := httptest.NewRequest("POST", "/submit", strings.NewReader(data))
        req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

        // We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
        rr := httptest.NewRecorder()
        handler(rr, req)

        // Check the status code is what we expect.
        if status := rr.Code; status != http.StatusOK {
                t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
        }

        // Check the response body is what we expect.
        expected := "Name = John\nEmail = john@example.com"
        if rr.Body.String() != expected {
                t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
        }
}
```

In this test, we create a new request with POST data and headers, then use the
`httptest.NewRecorder()` to record responses. We can then check if the response code
is as expected (200 OK), and also compare the body of the response to our expected
string.

Please note that this test assumes you have a function named 'handler' in your main
package which handles POST requests for '/submit'. If not, you need to import it from
where it's defined.

>>> it seems that you are missing a newline in your variable `expected`

You're absolutely right. I apologize for the confusion. The expected string should include a newline character ("\n") to match the format of the actual response from our
handler function. Here is the corrected test:

```go
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
```
This test should now pass as expected.
