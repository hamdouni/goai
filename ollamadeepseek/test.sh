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
