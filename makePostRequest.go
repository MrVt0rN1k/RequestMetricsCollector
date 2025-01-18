package main

import (
	"sync"
	"time"
)

func makePostRequest(url string, jsonData string, wg *sync.WaitGroup) {
	defer wg.Done()

	startTime := time.Now()

	makeRequest()
	fileSize(response)
	statusCode(response)
	lookEror(req)

	mutex.Lock()
	responseCounter++
	endTime := time.Now()

	countTime(endTime, startTime, responseCounter)
	switchTime(responseTime)
	mutex.Unlock()
}
