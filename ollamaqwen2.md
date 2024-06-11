# Déroulé du test avec le model 7b

>>> How to write a program in Go that handles an http POST request with multiple parameters ?

To create a Go program that accepts HTTP POST requests and processes multiple parameters, you can use the `net/http` package
provided by Go's standard library.

Here is a simple example:

```go
package main

import (
    "fmt"
    "log"
    "net/http"

    // This line imports the url.Values type for handling form data.
    _ "net/url"
)

// This function will handle the POST request and process its parameters
func handlePost(w http.ResponseWriter, r *http.Request) {
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
        params = append(params, k + "=" + vs[0])
    }
    // We can process the `params` array here if necessary.
}

func main() {
    http.HandleFunc("/", handlePost) // Set our handler function to be called when `/` is requested.

    log.Println("Server starting on port 8080")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal(err)
    }
}
```

In this program:

- `handlePost` handles an incoming HTTP POST request, parses the form parameters using `r.ParseForm()`, and writes the processed
information back to the client.

- The main function sets up a server that listens on port 8080 for incoming requests.

To run your Go application:
1. Save this code in a file named `main.go` or similar.
2. In the command line, navigate to the directory where you saved the file.
3. Run `go run main.go`.

This setup allows processing of form data sent via HTTP POST request by parsing the values and echoing them back.

Please note that the above example assumes that your `handlePost` function processes the parameters as needed for your specific
application.

Also, remember to install Go on your system before running this script using `go run main.go`.


>>> Can you provide some tests for this program ?

Sure! Below is a simple test suite that uses the `testing` package (`tgo`) provided by Go's standard library.

This suite tests different scenarios:

1. A successful POST request with all fields filled.
2. A POST request without any parameters.
3. A GET request instead of a POST.

Let's create a testing file named `test_main.go`.

```go
package main_test

import (
        "fmt"
        "net/http/httptest"
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

        t.Run("No fields filled in POST request", func(t *testing.T){
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

        t.Run("GET request", func(t *testing.T){
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

func HandlePost(w http.ResponseWriter, r *http.Request) { ... }
```
To run your tests:

1. Save this code into a new file named `test_main.go`.
2. Open your terminal or command line interface.
3. Navigate to the directory where you saved both your main (`main.go`) and test files (`test_main.go`).
4. Run `go test`. Go will automatically look for files with `.go` extension having function names starting with "Test" inside
them.

This setup uses an in-memory HTTP server created by `httptest.NewServer()` to simulate server responses, and then compares the
received response body and status code against expected results using `t.Errorf()`.

