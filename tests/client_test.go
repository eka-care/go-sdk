package records_test

import (
	"testing"

	"github.com/ekacare/go-sdk/client"
)

func TestNewClient(t *testing.T) {
	token := "api_key"
	client := client.NewClient("https://api.example.com", &token)
	if client.BaseURL != "https://api.example.com" {
		t.Errorf("Expected BaseURL to be 'https://api.example.com', got %s", client.BaseURL)
	}

}
