package http

import (
	"bytes"
	"net/http"
	"time"
)

// JSONRequest Sends a JSON request
func JSONRequest(target string, method string, data []byte, timeout time.Duration) (*http.Response, error) {
	client := &http.Client{
		Timeout: timeout,
	}
	req, err := http.NewRequest(method, target, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	res, err := client.Do(req)
	return res, err
}
