package http

import (
	"bytes"
	"io"
	"net/http"
)

type HttpClient struct {
	client *http.Client
}

func NewHttpClient() *HttpClient {
	return &HttpClient{
		client: &http.Client{},
	}
}

// ExecuteRequest Customised Request execute
func (h *HttpClient) ExecuteRequest(req *http.Request) (*http.Response, error) {
	res, err := h.client.Do(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// NewRequest creates a new customised HTTP request with the specified method, URL, headers, and body.
func (h *HttpClient) NewRequest(method, url string, headers map[string]string, body []byte) (*http.Request, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	if body != nil {
		req.Body = io.NopCloser(bytes.NewBuffer(body))
	}
	return req, nil
}

func (h *HttpClient) ExecuteCustomRequest(method, url string, headers map[string]string) (*http.Response, error) {
	req, err := h.NewRequest(method, url, headers, nil)
	if err != nil {
		return nil, err
	}
	res, err := h.ExecuteRequest(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
