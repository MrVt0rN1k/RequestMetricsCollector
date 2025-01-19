package main

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
)

func processResponse(resp *http.Response, body []byte) error {
	contentLength := resp.Header.Get("Content-Length")
	length, err := strconv.Atoi(contentLength)
	if err != nil {
		return fmt.Errorf("error converting Content-Length to integer: %v", err)
	}
	fileSize(length)

	lookFor := "code"
	refString := string(body)
	pattern := regexp.MustCompile(`\b` + regexp.QuoteMeta(lookFor) + `\b`)
	if pattern.MatchString(refString) {
		countError++
	}

	statusCode(resp)
	return nil
}
