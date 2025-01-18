package main

import (
	"fmt"
	"time"
)

func countTime(endTime time.Time, startTime time.Time, responseCounter int) time.Duration {
	responseTime := endTime.Sub(startTime)
	allResponceTime += responseTime
	avarageTime = allResponceTime / time.Duration(responseCounter)
	fmt.Printf("\r\033[1;36mWorking time: %v ", allResponceTime/time.Duration(requestsPerSecond))
	return responseTime
}
