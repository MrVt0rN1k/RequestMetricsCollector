package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"sync"
	"time"

	"github.com/gorilla/websocket"
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
	allSizeFile, less1mb, between1mband5mb, between5mband10mb, more10bm, size0, size int
	allResponceTime                                                                  time.Duration
	avarageTime                                                                      time.Duration
	mutex                                                                            sync.Mutex
	jsonData                                                                         string
	url                                                                              string
	duration                                                                         int
	requestsPerSecond                                                                int
	block                                                                            int
)

func makePostRequest(url string, jsonData string, wg *sync.WaitGroup) {
	defer wg.Done()

	startTime := time.Now()

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(jsonData)))
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		return
	}
	body, err := ioutil.ReadAll(req.Body)
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
	size = length / 1024
	if err != nil {
		fmt.Println("Error converting Content-Length to integer:", err)
		return
	}

	switch {
	case size > 0 && size <= 1000:
		less1mb++
	case size > 1000 && size <= 5000:
		between1mband5mb++
	case size > 5000 && size <= 10000:
		between5mband10mb++
	case size >= 10000:
		more10bm++
	default:
		size0++
	}

	lookFor := "code"
	refString := string(body)
	pattern := regexp.MustCompile(`\b` + regexp.QuoteMeta(lookFor) + `\b`)
	if pattern.MatchString(refString) {
		countError++
	}

	switch {
	case resp.StatusCode == 200:
		counter200++
	case resp.StatusCode == 500:
		counter500++
	case resp.StatusCode == 502:
		counter502++
	case resp.StatusCode == 503:
		counter503++
	case resp.StatusCode == 504:
		counter504++
	default:
		counterOther++
	}
	endTime := time.Now()
	responseTime := endTime.Sub(startTime)
	allResponceTime += responseTime
	mutex.Lock()
	responseCounter++
	avarageTime = allResponceTime / time.Duration(responseCounter)
	fmt.Printf("\r\033[1;36mWorking time: %v ", allResponceTime)

	allSizeFile += size
	switch {
	case responseTime < 100*time.Millisecond:
		less100ms++
	case responseTime >= 100*time.Millisecond && responseTime < 200*time.Millisecond:
		between100and200ms++
	case responseTime >= 200*time.Millisecond && responseTime < 300*time.Millisecond:
		between200and300ms++
	case responseTime >= 300*time.Millisecond && responseTime < 400*time.Millisecond:
		between300and400ms++
	case responseTime >= 400*time.Millisecond && responseTime < 500*time.Millisecond:
		between400and500ms++
	case responseTime >= 500*time.Millisecond && responseTime < 1*time.Second:
		between500and1000ms++
	case responseTime >= 1*time.Second && responseTime < 2*time.Second:
		between1000and2000ms++
	case responseTime >= 2*time.Second && responseTime < 3*time.Second:
		between2000and3000ms++
	case responseTime >= 3*time.Second && responseTime < 4*time.Second:
		between3000and4000ms++
	case responseTime >= 4*time.Second && responseTime < 5*time.Second:
		between4000and5000ms++
	case responseTime >= 5*time.Second && responseTime < 6*time.Second:
		between5000and6000ms++
	case responseTime >= 6*time.Second && responseTime < 7*time.Second:
		between6000and7000ms++
	case responseTime >= 7*time.Second && responseTime < 8*time.Second:
		between7000and8000ms++
	case responseTime >= 8*time.Second && responseTime < 9*time.Second:
		between8000and9000ms++
	case responseTime >= 9*time.Second && responseTime < 10*time.Second:
		between9000and10000ms++
	case responseTime >= 10*time.Second && responseTime < 11*time.Second:
		between10000and11000ms++
	case responseTime >= 11*time.Second && responseTime < 12*time.Second:
		between11000and12000ms++
	case responseTime >= 12*time.Second && responseTime < 13*time.Second:
		between12000and13000ms++
	case responseTime >= 13*time.Second && responseTime < 14*time.Second:
		between13000and14000ms++
	case responseTime >= 14*time.Second && responseTime < 15*time.Second:
		between14000and15000ms++
	case responseTime >= 15*time.Second && responseTime < 16*time.Second:
		between15000and16000ms++
	case responseTime >= 16*time.Second && responseTime < 17*time.Second:
		between16000and17000ms++
	case responseTime >= 17*time.Second && responseTime < 18*time.Second:
		between17000and18000ms++
	case responseTime >= 18*time.Second && responseTime < 19*time.Second:
		between18000and19000ms++
	case responseTime >= 19*time.Second && responseTime < 20*time.Second:
		between19000and20000ms++
	default:
		morethan20000ms++
	}
	mutex.Unlock()
}

func makeWssRequest(url string, jsonData string, wg *sync.WaitGroup) {
	dialer := websocket.DefaultDialer

	conn, _, err := dialer.Dial(url, nil)
	if err != nil {
		fmt.Println("Error on connection:", err)
	}

	err = conn.WriteJSON(jsonData)
	if err != nil {
		fmt.Println("Error on send request:", err)
		return
	}

	err = conn.ReadJSON(&jsonData)
	if err != nil {
		fmt.Println("Error on read answer:", err)
		return
	}
	fmt.Printf("%+v\n", jsonData)
	defer conn.Close()
}
func main() {
	flag.StringVar(&jsonData, "d", `{"jsonrpc": "2.0","method": "eth_blockNumber","params": [],"id": "getblock.io"}`, "Request you want to use")
	flag.StringVar(&url, "u", "https://go.getblock.io/c0f16139fdfe439eaab426f21dc94e70", "url")
	flag.IntVar(&duration, "t", 1, "Lead Time")
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
	duration := time.Duration(duration) * time.Minute
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
	fmt.Printf("\033[33mBetween 1000 and 2000 ms: %d\n", between1000and2000ms)
	fmt.Printf("\033[33mBetween 2000 and 3000 ms: %d\n", between2000and3000ms)
	fmt.Printf("\033[33mBetween 3000 and 4000 ms: %d\n", between3000and4000ms)
	fmt.Printf("\033[33mBetween 4000 and 5000 ms: %d\n", between4000and5000ms)
	fmt.Printf("\033[33mBetween 5000 and 6000 ms: %d\n", between5000and6000ms)
	fmt.Printf("\033[33mBetween 6000 and 7000 ms: %d\n", between6000and7000ms)
	fmt.Printf("\033[33mBetween 7000 and 8000 ms: %d\n", between7000and8000ms)
	fmt.Printf("\033[33mBetween 8000 and 9000 ms: %d\n", between8000and9000ms)
	fmt.Printf("\033[33mBetween 9000 and 10000 ms: %d\n", between9000and10000ms)
	fmt.Printf("\033[33mBetween 10000 and 11000 ms: %d\n", between10000and11000ms)
	fmt.Printf("\033[33mBetween 11000 and 12000 ms: %d\n", between11000and12000ms)
	fmt.Printf("\033[33mBetween 12000 and 13000 ms: %d\n", between12000and13000ms)
	fmt.Printf("\033[33mBetween 13000 and 14000 ms: %d\n", between13000and14000ms)
	fmt.Printf("\033[33mBetween 14000 and 15000 ms: %d\n", between14000and15000ms)
	fmt.Printf("\033[33mBetween 15000 and 16000 ms: %d\n", between15000and16000ms)
	fmt.Printf("\033[33mBetween 16000 and 17000 ms: %d\n", between16000and17000ms)
	fmt.Printf("\033[33mBetween 17000 and 18000 ms: %d\n", between17000and18000ms)
	fmt.Printf("\033[33mBetween 18000 and 19000 ms: %d\n", between18000and19000ms)
	fmt.Printf("\033[33mBetween 19000 and 20000 ms: %d\n", between19000and20000ms)
	fmt.Printf("\033[31mMore than 20000 ms: %d\n", morethan20000ms)
	fmt.Printf("\033[32mAvarage response time: %v\n", avarageTime)
	fmt.Printf("\033[32mLess than 1mb: %d\n", less1mb)
	fmt.Printf("\033[33mBetween 1mb and 5mb: %d\n", between1mband5mb)
	fmt.Printf("\033[33mBetween 5mb and 10mb: %d\n", between5mband10mb)
	fmt.Printf("\033[31mMore than 10mb: %d\n", more10bm)
	fmt.Printf("\033[32mAvarage size file: %d byte\n", allSizeFile/responseCounter)
}
