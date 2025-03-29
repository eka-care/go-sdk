package records

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"sync"
	"time"
)

func getStringAsPointer(s string) *string {
	return &s
}

func (r *RecordsService) getUploadURL() string {
	return fmt.Sprintf("%s%s", r.client.GetBaseURL(), UploadRecordAPIPath)
}

func (r *RecordsService) getAuthorization(batchRequest []BatchRequest) (*authorizationResponse, error) {
	data, _ := json.Marshal(map[string]interface{}{"batch_request": batchRequest})
	resp, err := r.client.Request(http.MethodPost, r.getUploadURL(), data, nil)
	if err != nil {
		return nil, err
	}

	var authResponse authorizationResponse
	if err = json.NewDecoder(resp.Body).Decode(&authResponse); err != nil {
		return nil, err
	}

	return &authResponse, nil
}

func (r *RecordsService) UploadDocument(batchRequest []BatchRequest) (*UploadResponse, error) {
	authResp, err := r.getAuthorization(batchRequest)
	if err != nil {
		return nil, err
	}

	if len(authResp.BatchResponse) == 0 {
		return nil, errors.New("no upload URL received")
	}

	var response UploadResponse
	for i, batch := range authResp.BatchResponse {
		if len(batch.Forms) == 0 {
			response.DocumentIDs = append(response.DocumentIDs, nil)
			continue
		}

		var errUpload error
		var wg sync.WaitGroup
		for j := range batch.Forms {
			wg.Add(1)
			go func(k int) {
				defer wg.Done()
				if err := r.upload(batchRequest[i].Files[k].Content, batch.Forms[k].URL, batch.Forms[k].Fields); err != nil {
					errUpload = err
					return
				}
			}(j)
		}

		wg.Wait()

		if errUpload != nil {
			response.DocumentIDs = append(response.DocumentIDs, nil)
			continue
		}

		response.DocumentIDs = append(response.DocumentIDs, getStringAsPointer(batch.DocumentID))
	}

	return &response, nil
}

func (r *RecordsService) upload(file []byte, url string, fields map[string]string) error {
	var b bytes.Buffer
	writer := multipart.NewWriter(&b)

	// Write form fields
	for key, val := range fields {
		if err := writer.WriteField(key, val); err != nil {
			return err
		}
	}

	// Create form file field
	part, err := writer.CreateFormFile("file", "randomfilename")
	if err != nil {
		return err
	}

	// Copy file data
	_, err = part.Write(file)
	if err != nil {
		return fmt.Errorf("error writing file to multipart: %v", err)
	}

	// Close writer to finalize multipart body
	writer.Close()

	// Create HTTP request with Content-Length
	req, err := http.NewRequest(http.MethodPost, url, &b)
	if err != nil {
		return err
	}
	req.Header.Set(HeaderContentType, writer.FormDataContentType())

	// Use an HTTP client with a timeout
	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode >= 300 {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("upload failed: %s", string(body))
	}

	return nil
}
