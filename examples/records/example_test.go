package records_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/ekacare/go-sdk/client"
	"github.com/ekacare/go-sdk/records"
)

func TestUploadSingleLapReport(t *testing.T) {
	token := ""
	host := ""
	client := client.NewClient(host, &token, nil, nil, nil)

	recordsService := records.NewRecordsService(client)
	filePath := ""
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
	token := ""
	host := ""
	client := client.NewClient(host, &token, nil, nil, nil)

	recordsService := records.NewRecordsService(client)
	filePath := ""
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
