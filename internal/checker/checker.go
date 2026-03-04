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

func CheckSite(url string) Result {
	client := http.Client{Timeout: 15 * time.Second}
	res, err := client.Get(url)

	if err != nil {
		return Result{URL: url, Error: err}
	}
	errBodyClose := res.Body.Close()
	if errBodyClose != nil {
		return Result{URL: url, Error: errBodyClose}
	}

	return Result{
		URL:    url,
		Code:   res.StatusCode,
		Status: res.Status,
		Error:  nil,
	}
}
