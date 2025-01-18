package main

import "net/http"

func statusCode(resp *http.Response) {
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
}
