package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func makeRequest() (*http.Request, *http.Response) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(jsonData)))
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept-Encoding", "gzip")
	client := &http.Client{}

	response, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error making request: %v\n", err)
	}
	return req, response
}
