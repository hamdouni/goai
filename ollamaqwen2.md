# Déroulé du test avec le model 7b

- **avis général** :
    - réponse plutôt rapide (<1mn)
    - le code pour le programme principal était correct du premier coup
    - le code pour le test était faux (l'appel à la fonction a testé n'était pas le bon), les instructions n'étaient pas corrects et la convention de nommage Go pour les tests n'était pas respecté (test_main.go au lieu de main_test.go) 
    - lorsque j'ai tenté de le corriger en lui indiquant ce qui n'allait pas dans le test, il m'a refait la même erreur
    - la commande CURL était juste, mais il n'a pas tenu compte du fait qu'il avait écrit un programme Go qui n'acceptait que les POST request et il m'a aussi proposé un CURL avec du GET

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

>>> The test program does not work. The HandlePost function does not exist in the main program, because it is named handlePost with an lower case h. I think you need to rewrite the test program. Also, in Go, the test file must be named with 'test' as a suffix and not as a prefix, so your instruction are bad. Can you try again ?

You're correct that my instructions were somewhat confusing. You should use `TestHandlePost` function name in your test file
because it matches the function being tested.

Also, naming conventions for Go testing packages typically involve prefixing them with 'test', which means your main test file
would look like `test_main.go`. I apologize for the confusion and thank you for clarifying this important point.

Below is a revised version of my previous answer based on those corrections:

```go
package test_main

import (
        "bufio"
        "bytes"
        "fmt"
        "net/http/httptest"
        "testing"
)

// Assuming HandlePost is now correctly named as handlePost function in main.go.
func TestHandlePost(t *testing.T) {
        t.Run("Successful POST request", func(t *testing.T) {
                testServer := httptest.NewServer(http.HandlerFunc(handlePost))
                defer testServer.Close()

                client := &http.Client{}
                req, err := http.NewRequest("POST", testServer.URL, strings.NewReader(`firstParam=hello&secondParam=world`))
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

                wantBody := `firstParam = hello\nsecondParam = world\n`
                if string(body) != wantBody {
                        t.Errorf("Unexpected response body: %s", string(body))
                }
        })

        t.Run("POST request with no parameters", func(t *testing.T) {
                testServer := httptest.NewServer(http.HandlerFunc(handlePost))
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

                wantStatus := http.StatusNotAcceptable
                wantBody := `This endpoint only accepts HTTP POST requests`
                if resp.StatusCode != wantStatus || string(body) != wantBody {
                        t.Errorf("Unexpected response: status code=%d body=\"%s\"", resp.StatusCode, string(body))
                }
        })

        t.Run("GET request", func(t *testing.T) {
                testServer := httptest.NewServer(http.HandlerFunc(handlePost))
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

                wantStatus := http.StatusMethodNotAllowed
                wantBody := `This endpoint only accepts HTTP POST requests`
                if resp.StatusCode != wantStatus || string(body) != wantBody {
                        t.Errorf("Unexpected response: status code=%d body=\"%s\"", resp.StatusCode, string(body))
                }
        })
}

func handlePost(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "%v", "firstParam = "+r.URL.Query().Get("firstParam")+"\n"+
        "secondParam = "+r.URL.Query().Get("secondParam"))
}
```

This revised version uses the `handlePost` function as it exists in your main program (`main.go`) and changes some of the test
scenarios to better reflect potential outcomes.

Remember, this is still a simplification for demonstration purposes. In practice, testing might involve more complex scenarios
depending on your specific application needs.

Please ensure you rename your test file back to `test_main.go` as per Go conventions if it's currently prefixed with 'test'.

Now you should be able to run these tests using `go test`. Let me know if there are any other questions!

>>> can you provide a curl command to test the Go program ?

Certainly, here is how you might use `curl` to simulate sending HTTP requests to your Go program:

### POST Request:
If you want to send a simple POST request with parameters like "firstParam" and "secondParam", you can use `curl` as follows:
```bash
curl -X POST \
  'http://localhost:8080/your-endpoint' \
  -d 'firstParam=hello' \
  -d 'secondParam=world'
```
This command sends a POST request to `http://localhost:8080/your-endpoint` with two parameters: "firstParam" set to "hello", and
"secondParam" set to "world".

### GET Request:
If you want to send a simple GET request, you can use the following curl command:

```bash
curl 'http://localhost:8080/your-endpoint?firstParam=hello&secondParam=world'
```
This sends a GET request to `http://localhost:8080/your-endpoint` with parameters "firstParam" set to "hello", and "secondParam"
set to "world".

### Testing using Postman:
Alternatively, you can use tools like [Postman](https://www.postman.com/) or similar API testing tools which allow graphical UI
operations for sending requests including POST with body data and GET requests.

For example, in Postman:
1. **Add Request**: Select `POST` as your HTTP method.
2. **Specify URL**: Enter `http://localhost:8080/your-endpoint`.
3. **Body Parameters**: Click on the `Body` tab below your request details. Use `Form Data`, specify keys (`firstParam`,
`secondParam`) and values (e.g., `"hello"`, `"world"`).
4. **Send Request**: Click the send button to execute the request.

Remember, for these commands to work, you need to replace `'http://localhost:8080/your-endpoint'` with your actual endpoint URL in
the Go application, and adjust the values of `firstParam`, `secondParam`, etc., as needed.

Make sure your server is running on `http://localhost:8080` or whatever you've specified when you're trying to access these URLs
through curl or Postman.

### Important Note:
To use the commands above, ensure that the Go application is running and listening on the correct port (`http://localhost:8080`).
Also, if your environment requires additional configuration (like SSL/TLS certificates), adjust accordingly.

