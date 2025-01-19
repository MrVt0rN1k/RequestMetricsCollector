package main

import (
	"fmt"
	"time"
)

func updateStats(startTime time.Time) {
	endTime := time.Now()
	responseTime := endTime.Sub(startTime)
	mutex.Lock()
	defer mutex.Unlock()
	allResponceTime += responseTime
	responseCounter++
	avarageTime = allResponceTime / time.Duration(responseCounter)
	fmt.Printf("\r\033[1;36mWorking time: %v ", allResponceTime/time.Duration(requestsPerSecond))
	allSizeFile += size
	switchTime(responseTime)
}
