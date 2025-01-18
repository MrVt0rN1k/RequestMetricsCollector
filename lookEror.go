package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
)

func lookEror(req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		fmt.Printf("client: error reading response body: %s\n", err)
		return
	}

	lookFor := "code"
	refString := string(body)
	pattern := regexp.MustCompile(`\b` + regexp.QuoteMeta(lookFor) + `\b`)
	if pattern.MatchString(refString) {
		countError++
	}
}
