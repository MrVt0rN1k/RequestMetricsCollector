package main

import (
	"flag"
	"fmt"
	"net/http"
	"sync"
	"time"
)

var (
	responseCounter                                            int
	counterOther                                               int
	countError                                                 int
	counter200, counter500, counter502, counter503, counter504 int
	less100ms, between100and200ms, between200and300ms, between300and400ms, between400and500ms, between500and1000ms, between1000and2000ms,
	between2000and3000ms, between3000and4000ms, between4000and5000ms, between5000and6000ms, between6000and7000ms, between7000and8000ms,
	between8000and9000ms, between9000and10000ms, between10000and11000ms, between11000and12000ms, between12000and13000ms, between13000and14000ms,
	between14000and15000ms, between15000and16000ms, between16000and17000ms, between17000and18000ms, between18000and19000ms, between19000and20000ms,
	morethan20000ms int
	size, size0, allSizeFile, less1mb, between1mband5mb, between5mband10mb, more10bm int
	allResponceTime                                                                  time.Duration
	avarageTime                                                                      time.Duration
	responseTime                                                                     time.Duration
	mutex                                                                            sync.Mutex
	jsonData                                                                         string
	url                                                                              string
	duration                                                                         int
	requestsPerSecond                                                                int
	block                                                                            int
	req                                                                              *http.Request
	resp                                                                             *http.Response
)

func main() {
	flag.StringVar(&jsonData, "d", `{"jsonrpc": "2.0","method": "eth_blockNumber","params": [],"id": "getblock.io"}`, "Request you want to use")
	flag.StringVar(&url, "u", "https://go.getblock.io/c0f16139fdfe439eaab426f21dc94e70", "url")
	flag.IntVar(&duration, "t", 30, "Lead Time")
	flag.IntVar(&requestsPerSecond, "r", 5, "RPS")
	flag.IntVar(&block, "b", 0, "Block number for methed get_block on sol")
	flag.Parse()

	if block != 0 {

	}
	counterOther = 0
	counter200 = 0
	counter500 = 0
	counter502 = 0
	counter503 = 0
	counter504 = 0
	allSizeFile = 0
	var wg sync.WaitGroup
	duration := time.Duration(duration) * time.Second
	startTime := time.Now()
	for time.Since(startTime) < duration {
		for i := 0; i < requestsPerSecond; i++ {
			wg.Add(1)
			go makePostRequest(url, jsonData, &wg)
		}
		time.Sleep(1 * time.Second)
		block++

	}

	wg.Wait()
	fmt.Println()
	fmt.Printf("\033[32mTotal HTTP responses: %d\n", responseCounter)
	fmt.Printf("\033[32mTotal error responses: %d\n", countError)
	fmt.Printf("\033[32mTotal HTTP Statuscode 200: %d\n", counter200)
	fmt.Printf("\033[31mTotal HTTP Statuscode 500: %d\n", counter500)
	fmt.Printf("\033[31mTotal HTTP Statuscode 502: %d\n", counter502)
	fmt.Printf("\033[31mTotal HTTP Statuscode 503: %d\n", counter503)
	fmt.Printf("\033[31mTotal HTTP Statuscode 504: %d\n", counter504)
	fmt.Printf("\033[31mTotal HTTP Statuscode Unknow: %d\n", counterOther)
	fmt.Printf("\033[32mLess than 100 ms: %d\n", less100ms)
	fmt.Printf("\033[32mBetween 100 and 200 ms: %d\n", between100and200ms)
	fmt.Printf("\033[32mBetween 200 and 300 ms: %d\n", between200and300ms)
	fmt.Printf("\033[32mBetween 300 and 400 ms: %d\n", between300and400ms)
	fmt.Printf("\033[32mBetween 400 and 500 ms: %d\n", between400and500ms)
	fmt.Printf("\033[32mBetween 500 and 1000 ms: %d\n", between500and1000ms)
	fmt.Printf("\033[31mBetween 1000 and 2000 ms: %d\n", between1000and2000ms)
	fmt.Printf("\033[31mBetween 2000 and 3000 ms: %d\n", between2000and3000ms)
	fmt.Printf("\033[31mBetween 3000 and 4000 ms: %d\n", between3000and4000ms)
	fmt.Printf("\033[31mBetween 4000 and 5000 ms: %d\n", between4000and5000ms)
	fmt.Printf("\033[31mBetween 5000 and 6000 ms: %d\n", between5000and6000ms)
	fmt.Printf("\033[31mBetween 6000 and 7000 ms: %d\n", between6000and7000ms)
	fmt.Printf("\033[31mBetween 7000 and 8000 ms: %d\n", between7000and8000ms)
	fmt.Printf("\033[31mBetween 8000 and 9000 ms: %d\n", between8000and9000ms)
	fmt.Printf("\033[31mBetween 9000 and 10000 ms: %d\n", between9000and10000ms)
	fmt.Printf("\033[31mBetween 10000 and 11000 ms: %d\n", between10000and11000ms)
	fmt.Printf("\033[31mBetween 11000 and 12000 ms: %d\n", between11000and12000ms)
	fmt.Printf("\033[31mBetween 12000 and 13000 ms: %d\n", between12000and13000ms)
	fmt.Printf("\033[31mBetween 13000 and 14000 ms: %d\n", between13000and14000ms)
	fmt.Printf("\033[31mBetween 14000 and 15000 ms: %d\n", between14000and15000ms)
	fmt.Printf("\033[31mBetween 15000 and 16000 ms: %d\n", between15000and16000ms)
	fmt.Printf("\033[31mBetween 16000 and 17000 ms: %d\n", between16000and17000ms)
	fmt.Printf("\033[31mBetween 17000 and 18000 ms: %d\n", between17000and18000ms)
	fmt.Printf("\033[31mBetween 18000 and 19000 ms: %d\n", between18000and19000ms)
	fmt.Printf("\033[31mBetween 19000 and 20000 ms: %d\n", between19000and20000ms)
	fmt.Printf("\033[31mMore than 20000 ms: %d\n", morethan20000ms)
	fmt.Printf("\033[32mAvarage response time: %v\n", avarageTime)
	fmt.Printf("\033[32mLess than 1mb: %d\n", less1mb)
	fmt.Printf("\033[33mBetween 1mb and 5mb: %d\n", between1mband5mb)
	fmt.Printf("\033[33mBetween 5mb and 10mb: %d\n", between5mband10mb)
	fmt.Printf("\033[31mMore than 10mb: %d\n", more10bm)
	fmt.Printf("\033[32mAvarage size file: %d byte\n", allSizeFile/responseCounter)
}
