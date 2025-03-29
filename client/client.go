package client

import (
	"net/http"
)

type ClientInterface interface {
	Request(method, url string, data []byte, headers map[string]string) (*http.Response, error)
	GetBaseURL() string
}

type Client struct {
	baseURL    string
	token      string
	HTTPClient *http.Client
}

func NewClient(baseURL, token string) ClientInterface {
	return &Client{
		baseURL:    baseURL,
		token:      token,
		HTTPClient: &http.Client{},
	}
}
