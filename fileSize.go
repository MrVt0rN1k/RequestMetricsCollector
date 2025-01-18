package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func fileSize(resp *http.Response) {
	defer resp.Body.Close()
	contentLength := resp.Header.Get("Content-Length")
	length, err := strconv.Atoi(contentLength)
	if length < 1024 {
		size = length
	} else {
		size = length / 1024
	}
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
	allSizeFile += size
}
