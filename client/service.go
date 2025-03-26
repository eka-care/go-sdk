package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

func (c *Client) Request(method, url string, data []byte, headers map[string]string) (*http.Response, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", *c.token))
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

func (c *Client) Authenticate() error {
	data := map[string]string{
		"client_id":     *c.id,
		"client_secret": *c.secret,
	}
	if c.key != nil && *c.key != "" {
		data["api_key"] = *c.key
	}

	body, _ := json.Marshal(data)
	resp, err := c.Request("POST", c.url, body, nil)

	if err != nil {
		return err
	}

	var authResponse authResponse
	err = json.NewDecoder(resp.Body).Decode(&authResponse)
	if err != nil {
		return err
	}

	c.token = &authResponse.AccessToken
	return nil
}
