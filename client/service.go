package client

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
)

const AuthorizationHeader string = "Authorization"
const BearerPrefix string = "Bearer "

func (c *Client) Request(method, url string, data []byte, headers map[string]string) (*http.Response, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set(AuthorizationHeader, fmt.Sprintf("%s%v", BearerPrefix, c.token))
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 400 {
		return nil, errors.New("API request failed")
	}

	return resp, nil
}

func (c *Client) GetBaseURL() string {
	return c.baseURL
}
