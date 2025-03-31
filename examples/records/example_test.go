package records_test

import (
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

	// Get file size using Stat
	fileInfo, err := file.Stat()
	if err != nil {
		t.Fatalf("failed to get file info: %+v", err)
	}
	fileSize := fileInfo.Size()

	var batchRequest []records.BatchRequest
	batchRequest = append(batchRequest, records.BatchRequest{
		DocumentType: records.LabReportQP,
		Files: []records.File{
			{
				Content:     file,
				ContentType: records.GetContentType(file),
				FileSize:    fileSize,
			},
		},
	})

	documentIDs, err := recordsService.UploadDocument(records.UploadRequest{Request: batchRequest, Tasks: []records.Task{records.SmartReportTaskQP, records.PIITaskQP}})
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

	// Get file size using Stat
	fileInfo, err := file.Stat()
	if err != nil {
		t.Fatalf("failed to get file info: %+v", err)
	}
	fileSize := fileInfo.Size()

	var batchRequest []records.BatchRequest
	batchRequest = append(batchRequest, records.BatchRequest{
		DocumentType: records.LabReportQP,
		Files: []records.File{
			{Content: file, ContentType: records.GetContentType(file), FileSize: fileSize},
			{Content: file, ContentType: records.GetContentType(file), FileSize: fileSize},
		},
	})

	documentID, err := recordsService.UploadDocument(records.UploadRequest{Request: batchRequest})
	if err != nil {
		t.Fatalf("Expected success', got err: %+v", err)
	} else {
		t.Log("Uploaded Document ID: ", documentID)
	}
}
