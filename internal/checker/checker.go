package checker

import (
	"net/http"
	"time"
)

type Result struct {
	URL    string
	Code   int
	Status string
	Error  error
}

var client = &http.Client{Timeout: 10 * time.Second}

func CheckSite(url string) Result {
	resp, err := client.Get(url)

	if err != nil {
		return Result{URL: url, Error: err}
	}
	defer resp.Body.Close()

	return Result{
		URL:    url,
		Code:   resp.StatusCode,
		Status: resp.Status,
		Error:  nil,
	}
}
