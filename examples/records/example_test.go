package records_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/ekacare/go-sdk/client"
	"github.com/ekacare/go-sdk/records"
)

func TestUploadSingleLapReport(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJhZHNpbmRpYXN0YWdpbmciLCJjLWlkIjoiYWRzaW5kaWFzdGFnaW5nIiwiZXhwIjoxNzQyOTkxODQ4LCJpYXQiOjE3NDI5OTEyNDgsImlzcyI6ImVrYS5jYXJlIn0.0mGmcV44WhThxu7dWqNc8yc0yJriXDGUklnz-719gfY"
	client := client.NewClient("https://api.dev.eka.care", &token, nil, nil, nil)

	recordsService := records.NewRecordsService(client)
	filePath := "lab_report.jpg"
	data, err := os.Open(filePath)
	if err != nil {
		t.Errorf("file not found', got err: %s", err)
	}

	files := []records.FileRequest{
		{FileName: filePath, Reader: data},
	}
	documentID, err := recordsService.UploadDocument(files, "lr", nil, nil, nil)
	if err != nil {
		fmt.Println("Error:", err)
		t.Errorf("Expected success', got err: %s", err)
	} else {
		fmt.Println("Uploaded Document ID:", documentID)
	}
}

func TestUploadMultipleLapReport(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJhZHNpbmRpYXN0YWdpbmciLCJjLWlkIjoiYWRzaW5kaWFzdGFnaW5nIiwiZXhwIjoxNzQyOTkxODQ4LCJpYXQiOjE3NDI5OTEyNDgsImlzcyI6ImVrYS5jYXJlIn0.0mGmcV44WhThxu7dWqNc8yc0yJriXDGUklnz-719gfY"
	client := client.NewClient("https://api.dev.eka.care", &token, nil, nil, nil)

	recordsService := records.NewRecordsService(client)
	filePath := "lab_report.jpg"
	data, err := os.Open(filePath)
	if err != nil {
		t.Errorf("file not found', got err: %s", err)
	}

	files := []records.FileRequest{
		{FileName: filePath, Reader: data},
		{FileName: filePath, Reader: data},
	}

	documentID, err := recordsService.UploadDocument(files, "lr", nil, nil, nil)
	if err != nil {
		fmt.Println("Error:", err)
		t.Errorf("Expected success', got err: %s", err)
	} else {
		fmt.Println("Uploaded Document ID:", documentID)
	}
}
