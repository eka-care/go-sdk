package client

import (
	"fmt"
	"net/http"
)

type Client struct {
	BaseURL    string
	id         *string
	secret     *string
	key        *string
	token      *string
	url        string
	HTTPClient *http.Client
}

func NewClient(baseURL string, token, id, secret, key *string) *Client {
	url := fmt.Sprintf("%s%s", baseURL, AUTH_ENDPOINT)
	return &Client{
		BaseURL:    baseURL,
		id:         id,
		secret:     secret,
		key:        key,
		url:        url,
		token:      token,
		HTTPClient: &http.Client{},
	}
}
