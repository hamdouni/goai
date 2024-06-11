#!/bin/sh

# Send POST request
response=$(curl -s -d "name=John&email=john@example.com" -H "Content-Type: application/x-www-form-urlencoded" -X POST http://localhost:8080/submit)

echo "${#response} $response"

# Check response body
expected="Name = John
Email = john@example.com"
echo "${#expected} $expected"
if [ "$response" != "$expected" ]; then
    echo "Error: Unexpected response."
    exit 1
fi

