package records_test

import (
	"fmt"
	"testing"

	"github.com/ekacare/go-sdk/client"
	"github.com/ekacare/go-sdk/records"
)

func TestUploadLapReport(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJhZHNpbmRpYXN0YWdpbmciLCJjLWlkIjoiYWRzaW5kaWFzdGFnaW5nIiwiZXhwIjoxNzQyOTg2ODI1LCJpYXQiOjE3NDI5ODYyMjUsImlzcyI6ImVrYS5jYXJlIn0.To1u8VHC0kFWwpJyYddywmlv6nR7v02zQQIwT-iC-Y4"
	client := client.NewClient("https://api.dev.eka.care", &token, nil, nil, nil)

	recordsService := records.NewRecordsService(client)
	documentID, err := recordsService.UploadDocument("lab_report.jpg", "lr", nil, nil, nil)
	if err != nil {
		fmt.Println("Error:", err)
		t.Errorf("Expected success', got err: %s", err)
	} else {
		fmt.Println("Uploaded Document ID:", documentID)
	}
}
