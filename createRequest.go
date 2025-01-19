package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func createRequest(url, jsonData string) (*http.Request, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(jsonData)))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept-Encoding", "gzip")
	return req, nil
}
