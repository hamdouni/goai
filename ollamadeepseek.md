# Déroulé du test

>>> How to write a program in Go that handles an http POST request with multiple parameters ?
...

Here is a simple example of a HTTP server written in
Golang which can handle post requests from clients having
multiple parameters. This example uses the "net/http"
standard library for handling the HTTP protocol and
"encoding/json" for unmarshalling JSON payloads into Go
structs.

```go
package main

import (
        "encoding/json"
        "fmt"
        "log"
        "net/http"
)

type ParamsStruct struct {
        Param1 string `json:"param1"`
        Param2 int    `json:"param2"`
}

func handler(w http.ResponseWriter, r *http.Request) {
        // Decode the JSON request into a ParamsStruct
        var params ParamsStruct
        err := json.NewDecoder(r.Body).Decode(&params)
        if err != nil {
                http.Error(w, "Bad Request", http.StatusBadRequest)
                return
        }
        fmt.Printf("Received parameters: %v\n", params)
}

func main() {
        http.HandleFunc("/post-endpoint", handler)
        log.Fatal(http.ListenAndServe(":8080", nil))
}
```
In this example, the server is listening on port 8080 for
incoming POST requests at "/post-endpoint". The "handler"
function extracts the JSON payload from the request body
and unmarshals it into a struct called ParamsStruct.
After that you can handle your parameters as needed in
your code (e.g., by doing some calculations or data
processing).

Note: You need to install Go on your machine, set up
GOPATH correctly and run this program. The "http" package
provides the necessary functionality for setting up a
server and handling requests. After starting the server
with `go run main.go` (assuming you named your file as
'main.go'), you can send POST request to
`localhost:8080/post-endpoint` with JSON payload
containing parameters like {"param1": "value1", "param2":
2}.

>>> Can you provide some tests for this program ?

Sure, let's use the `net/http/httptest` package to test
our handler function. We'll create a new request and pass
it into the handler along with an httpRecorder which
mimics the HTTP response.

Here is how you can write tests for this program:

```go
package main

import (
        "bytes"
        "encoding/json"
        "net/http"
        "net/http/httptest"
        "testing"
)

func TestHandler(t *testing.T) {
        type testStruct struct {
                Param1 string `json:"param1"`
                Param2 int    `json:"param2"`
        }

        tcs := []struct {
                name       string
                body       testStruct
                wantStatus int
        }{
                {
                        "valid request",
                        testStruct{"value1", 2},
                        http.StatusOK,
                },
        }
        for _, tc := range tcs {
                t.Run(tc.name, func(t *testing.T) {
                        // Convert the test struct into a JSON string.
                        jsonStr, _ := json.Marshal(tc.body)

                        // Create a new request with JSON body.
                        req, err := http.NewRequest("POST", "/post-endpoint",
bytes.NewBufferString(string(jsonStr)))
                        if err != nil {
                                t.Fatalf("Could not create request: %v\n", err)
                        }

                        // Create a new ResponseRecorder which satisfies
http.ResponseWriter.
                        rec := httptest.NewRecorder()
                        handler(rec, req)

                        // Check the status code is what we expect.
                        if status := rec.Code; status != tc.wantStatus {
                                t.Errorf("Handler returned wrong status code: got %v
want %v", status, http.StatusOK)
                        }
                })
        }
}
```

In this test function, we create a slice of structs
(`tcs`), where each struct represents one test case. Each
struct contains the name of the test case, an instance of
`testStruct` to be sent as JSON in request body and the
expected HTTP status code.

Each test case is run by calling `http.NewRequest` with
method "POST", endpoint "/post-endpoint" and a JSON
string converted from `testStruct`. Then we call our
handler function passing this request and an instance of
httpTest.ResponseRecorder to capture the response. The
status code returned by the server is checked against the
expected value.

>>> Can you write tests with curl ?

Sure, here's how you can use `curl` command to test your
HTTP POST endpoint from a shell script:

Let's assume that you have started your Go program and it
is listening on port 8080. Here's an example of a bash
script (test.sh) that sends a curl request:

```bash
#!/bin/bash

# Define the endpoint
url="http://localhost:8080/post-endpoint"

# Define the JSON payload
json_payload='{"param1": "value1", "param2": 2}'

# Send POST request with curl
response=$(curl -s -X POST -H "Content-Type:
application/json" -d "$json_payload" $url)

# Check if the response is as expected
if [[ "$response" == '{"param1":"value1","param2":2}' ]];
then
    echo "Test passed."
else
    echo "Test failed. Expected:
'{\"param1\":\"value1\",\"param2\":2}', got: $response"
fi
```
You can run this script with `bash test.sh` command. It
sends a POST request to your server with JSON payload
{"param1": "value1", "param2": 2}, and then checks the
response. If it matches the expected output, it prints
"Test passed." Otherwise, it prints "Test failed." along
with the actual response from the server.

>>> the program does not seem to return a json response. Can yo
... u fix it ?

The provided code was returning an HTTP 200 (OK) status for all requests because it's only printing out the parameters in the console and not sending any JSON response back to the client. If we want to send a JSON
response, we need to write that as well. Here is how you can modify the handler function:

```go
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
```
In this code, we are sending a simple JSON response back to the client after successfully parsing and validating the request body. The response is written to `w` (which stands for HTTP writer), which sends it as a reply
to the client's request. The content-type of the response is set to "application/json". If there are any issues during JSON marshalling, an error response with status 500 is sent back to the client.


>>> Can you write the curl command to test this program ?

Certainly! Here's how you can use `curl` command to test your HTTP POST endpoint from a shell script (test.sh):

Let's assume that you have started your Go program and it is listening on port 8080. The following bash script sends a curl request:

```bash
#!/bin/bash

# Define the endpoint
url="http://localhost:8080/post-endpoint"

# Define the JSON payload
json_payload='{"param1": "value1", "param2": 2}'

# Send POST request with curl
response=$(curl -s -X POST -H "Content-Type: application/json" -d "$json_payload" $url)

# Check if the response is as expected
if [[ "$response" == '{"status":"success"}' ]]; then
    echo "Test passed."
else
    echo "Test failed. Expected: '{\"status\":\"success\"}', got: $response"
fi
```
You can run this script with `bash test.sh` command. It sends a POST request to your server with JSON payload {"param1": "value1", "param2": 2}, and then checks the response. If it matches the expected output, it prints
"Test passed." Otherwise, it prints "Test failed." along with the actual response from the server.


