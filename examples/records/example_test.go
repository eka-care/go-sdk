package records_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/ekacare/go-sdk/client"
	"github.com/ekacare/go-sdk/records"
)

func TestUploadSingleLapReport(t *testing.T) {
	token := os.Getenv("API_TOKEN")
	host := os.Getenv("API_HOST")
	if token == "" || host == "" {
		t.Errorf("Expected both token and host', got token: %s host %s", token, host)
	}

	client := client.NewClient(host, &token)
	recordsService := records.NewRecordsService(client)
	filePath := "lap_report.jpg"
	data, err := os.Open(filePath)
	if err != nil {
		t.Errorf("file not found', got err: %s", err)
	}

	files := []records.FileRequest{
		{
			FileName:     filePath,
			Reader:       data,
			DocumentType: records.LabReportQP,
		},
	}

	var batchRequest []records.BatchRequest
	for _, file := range files {
		fileSize := records.GetFileSize(file.Reader)
		contentType := records.GetContentType(file.FileName)
		batchRequest = append(batchRequest, records.BatchRequest{
			DocumentType: file.DocumentType,
			DocumentDate: file.DocumentDate,
			Tags:         nil,
			Files:        []records.File{{ContentType: contentType, FileSize: fileSize}},
		})
	}

	documentID, err := recordsService.UploadDocument(files, batchRequest)
	if err != nil {
		fmt.Println("Error:", err)
		t.Errorf("Expected success', got err: %s", err)
	} else {
		fmt.Println("Uploaded Document ID:", documentID)
	}
}

func TestUploadMultipleLapReport(t *testing.T) {
	token := os.Getenv("API_TOKEN")
	host := os.Getenv("API_HOST")
	if token == "" || host == "" {
		t.Errorf("Expected both token and host', got token: %s host %s", token, host)
	}

	client := client.NewClient(host, &token)
	recordsService := records.NewRecordsService(client)
	filePath := "lap_report.jpg"
	data, err := os.Open(filePath)
	if err != nil {
		t.Errorf("file not found', got err: %s", err)
	}

	files := []records.FileRequest{
		{FileName: filePath, Reader: data},
		{FileName: filePath, Reader: data},
	}

	var batchRequest []records.BatchRequest
	for _, file := range files {
		fileSize := records.GetFileSize(file.Reader)
		contentType := records.GetContentType(file.FileName)
		batchRequest = append(batchRequest, records.BatchRequest{
			DocumentType: file.DocumentType,
			DocumentDate: file.DocumentDate,
			Tags:         nil,
			Files:        []records.File{{ContentType: contentType, FileSize: fileSize}},
		})
	}

	documentID, err := recordsService.UploadDocument(files, batchRequest)
	if err != nil {
		fmt.Println("Error:", err)
		t.Errorf("Expected success', got err: %s", err)
	} else {
		fmt.Println("Uploaded Document ID:", documentID)
	}
}
