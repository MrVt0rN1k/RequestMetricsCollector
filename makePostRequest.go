package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"sync"
	"time"
)

func makePostRequest(url string, jsonData string, wg *sync.WaitGroup) {
	defer wg.Done()
	startTime := time.Now()
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(jsonData)))
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		return
	}
	body, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Printf("client: error reading response body: %s\n", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept-Encoding", "gzip")
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error making request: %v\n", err)
		return
	}

	defer resp.Body.Close()
	contentLength := resp.Header.Get("Content-Length")

	length, err := strconv.Atoi(contentLength)
	if err != nil {
		fmt.Println("Error converting Content-Length to integer:", err)
		return
	}

	fileSize(length)

	lookFor := "code"
	refString := string(body)
	pattern := regexp.MustCompile(`\b` + regexp.QuoteMeta(lookFor) + `\b`)
	if pattern.MatchString(refString) {
		countError++
	}

	statusCode(resp)
	endTime := time.Now()
	responseTime := endTime.Sub(startTime)
	allResponceTime += responseTime
	mutex.Lock()
	responseCounter++
	avarageTime = allResponceTime / time.Duration(responseCounter)
	fmt.Printf("\r\033[1;36mWorking time: %v ", allResponceTime/time.Duration(requestsPerSecond))
	allSizeFile += size
	switchTime(responseTime)
	mutex.Unlock()
}
