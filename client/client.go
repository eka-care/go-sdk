package client

import (
	"net/http"
)

type Client struct {
	BaseURL    string
	token      *string
	HTTPClient *http.Client
}

func NewClient(baseURL string, token *string) *Client {
	return &Client{
		BaseURL:    baseURL,
		token:      token,
		HTTPClient: &http.Client{},
	}
}
