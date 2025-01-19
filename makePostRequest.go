package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

func makePostRequest(url string, jsonData string, wg *sync.WaitGroup) {
	defer wg.Done()
	startTime := time.Now()

	req, err := createRequest(url, jsonData)
	if err != nil {
		fmt.Println(err)
		return
	}

	body, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Printf("client: error reading response body: %s\n", err)
		return
	}

	client := &http.Client{}
	resp, err := executeRequest(client, req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	if err := processResponse(resp, body); err != nil {
		fmt.Println(err)
		return
	}

	updateStats(startTime)
}
