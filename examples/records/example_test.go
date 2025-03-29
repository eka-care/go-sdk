package records_test

import (
	"io"
	"os"
	"testing"

	"github.com/ekacare/go-sdk/client"
	"github.com/ekacare/go-sdk/records"
)

func TestUploadSingleFile(t *testing.T) {
	token := os.Getenv("API_TOKEN")
	host := os.Getenv("API_HOST")
	if token == "" || host == "" {
		t.Fatalf("Expected both token and host', got token: %s host %s", token, host)
	}

	client := client.NewClient(host, token)
	recordsService := records.NewRecordsService(client)

	filePath := "lap_report.jpg"
	file, err := os.Open(filePath)
	if err != nil {
		t.Fatalf("file not found', got err: %+v", err)
	}

	defer file.Close() // Always close file

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		t.Fatalf("error reading file: %+v", err)
	}

	files := []records.BatchRequest{
		{
			Files: []records.File{
				{ContentType: records.GetContentType(fileBytes), FileSize: len(fileBytes)},
			},
			DocumentType: records.LabReportQP,
		},
	}

	var batchRequest []records.BatchRequest
	for _, file := range files {
		batchRequest = append(batchRequest, records.BatchRequest{
			DocumentType: file.DocumentType,
			DocumentDate: file.DocumentDate,
			Tags:         nil,
			Files:        []records.File{{ContentType: file.Files[0].ContentType, FileSize: file.Files[0].FileSize}},
		})
	}

	documentIDs, err := recordsService.UploadDocument(batchRequest)
	if err != nil {
		t.Fatalf("Expected success', got err: %+v", err)
	} else {
		t.Log("Uploaded Document ID: ", documentIDs)
	}
}

func TestUploadMultipleFiles(t *testing.T) {
	token := os.Getenv("API_TOKEN")
	host := os.Getenv("API_HOST")
	if token == "" || host == "" {
		t.Fatalf("Expected both token and host', got token: %s host %s", token, host)
	}

	client := client.NewClient(host, token)
	recordsService := records.NewRecordsService(client)
	filePath := "lap_report.jpg"
	file, err := os.Open(filePath)
	if err != nil {
		t.Fatalf("file not found', got err: %+v", err)
	}

	defer file.Close() // Always close file

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		t.Fatalf("error reading file: %+v", err)
	}

	files := []records.BatchRequest{
		{
			Files: []records.File{
				{ContentType: records.GetContentType(fileBytes), FileSize: len(fileBytes)},
				{ContentType: records.GetContentType(fileBytes), FileSize: len(fileBytes)},
			},
			DocumentType: records.LabReportQP,
		},
	}

	var batchRequest []records.BatchRequest
	for _, file := range files {
		batchRequest = append(batchRequest, records.BatchRequest{
			DocumentType: file.DocumentType,
			DocumentDate: file.DocumentDate,
			Tags:         nil,
			Files: []records.File{
				{ContentType: file.Files[0].ContentType, FileSize: file.Files[0].FileSize},
				{ContentType: file.Files[1].ContentType, FileSize: file.Files[1].FileSize},
			},
		})
	}

	documentID, err := recordsService.UploadDocument(batchRequest)
	if err != nil {
		t.Fatalf("Expected success', got err: %+v", err)
	} else {
		t.Log("Uploaded Document ID: ", documentID)
	}
}
