package records_test

import (
	"io"
	"os"
	"testing"

	"github.com/ekacare/go-sdk/client"
	"github.com/ekacare/go-sdk/records"
)

func TestUploadSingleLapReport(t *testing.T) {
	token := os.Getenv("API_TOKEN")
	host := os.Getenv("API_HOST")
	if token == "" || host == "" {
		t.Fatalf("Expected both token and host', got token: %s host %s", token, host)
	}
	client := client.NewClient(host, &token)
	recordsService := records.NewRecordsService(client)

	filePath := "lap_report.jpg"
	file, err := os.Open(filePath)
	if err != nil {
		t.Fatalf("file not found', got err: %s", err)
	}
	defer file.Close() // Always close file

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		t.Fatalf("error reading file: %s", err)
	}

	files := []records.FileRequest{
		{
			FileName:     filePath,
			Content:      fileBytes,
			DocumentType: records.LabReportQP,
		},
	}

	var batchRequest []records.BatchRequest
	for _, file := range files {
		fileSize := len(file.Content)
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
		t.Fatalf("Expected success', got err: %s", err)
	} else {
		t.Log("Uploaded Document ID:", documentID)
	}
}

func TestUploadMultipleLapReport(t *testing.T) {
	token := os.Getenv("API_TOKEN")
	host := os.Getenv("API_HOST")
	if token == "" || host == "" {
		t.Fatalf("Expected both token and host', got token: %s host %s", token, host)
	}

	client := client.NewClient(host, &token)
	recordsService := records.NewRecordsService(client)
	filePath := "lap_report.jpg"
	file, err := os.Open(filePath)
	if err != nil {
		t.Fatalf("file not found', got err: %s", err)
	}

	defer file.Close() // Always close file

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		t.Fatalf("error reading file: %s", err)
	}

	files := []records.FileRequest{
		{FileName: filePath, Content: fileBytes},
		{FileName: filePath, Content: fileBytes},
	}

	var batchRequest []records.BatchRequest
	for _, file := range files {
		fileSize := len(file.Content)
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
		t.Fatalf("Expected success', got err: %s", err)
	} else {
		t.Log("Uploaded Document ID:", documentID)
	}
}
