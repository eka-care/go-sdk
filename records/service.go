package records

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"sync"
	"time"
)

func getStringAsPointer(s string) *string {
	return &s
}

func (r *RecordsService) getUploadURL() string {
	return fmt.Sprintf("%s%s", r.client.GetBaseURL(), UploadRecordAPIPath)
}

func (r *RecordsService) getAuthorization(uploadReq UploadRequest) (*authorizationResponse, error) {
	data, err := json.Marshal(map[string]interface{}{"batch_request": uploadReq.Request})
	if err != nil {
		return nil, fmt.Errorf("failed to marshal batch request: %w", err)
	}

	v := url.Values{}
	for _, task := range uploadReq.Tasks {
		v.Add("task", string(task))
	}

	if uploadReq.Batch {
		v.Add("batch", "true")
	}

	uploadURL := fmt.Sprintf("%s?%s", r.getUploadURL(), v.Encode())
	resp, err := r.client.Request(http.MethodPost, uploadURL, data, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get authorization: %w", err)
	}
	defer resp.Body.Close()

	var authResponse authorizationResponse
	if err = json.NewDecoder(resp.Body).Decode(&authResponse); err != nil {
		return nil, fmt.Errorf("failed to decode authorization response: %w", err)
	}

	return &authResponse, nil
}

func (r *RecordsService) UploadDocument(batchRequest UploadRequest) (*UploadResponse, error) {
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

		errChan := make(chan error, len(batch.Forms))

		var wg sync.WaitGroup
		for j := range batch.Forms {
			wg.Add(1)
			go func(k int) {
				defer wg.Done()
				fileReader := batchRequest.Request[i].Files[k].Content // io.Reader
				err := r.upload(fileReader, batch.Forms[k].URL, batch.Forms[k].Fields)
				if err != nil {
					errChan <- fmt.Errorf("failed to upload file %d in batch %d: %w", k, i, err)
					return
				}
				errChan <- nil
			}(j)
		}

		wg.Wait()
		close(errChan)

		var uploadError error
		for err := range errChan {
			if err != nil {
				uploadError = err
				break
			}
		}

		if uploadError != nil {
			response.DocumentIDs = append(response.DocumentIDs, nil)
		} else {
			response.DocumentIDs = append(response.DocumentIDs, getStringAsPointer(batch.DocumentID))
		}
	}

	return &response, nil
}

func (r *RecordsService) upload(file io.Reader, url string, fields map[string]string) error {
	var b bytes.Buffer
	writer := multipart.NewWriter(&b)

	// Write form fields
	for key, val := range fields {
		if err := writer.WriteField(key, val); err != nil {
			return fmt.Errorf("failed to write form field %s: %w", key, err)
		}
	}

	// Create form file field
	part, err := writer.CreateFormFile("file", "randomfilename")
	if err != nil {
		return fmt.Errorf("failed to create form file: %w", err)
	}

	// Copy file data
	if _, err = io.Copy(part, file); err != nil {
		return fmt.Errorf("failed to copy file data to multipart writer: %w", err)
	}

	// Close writer to finalize multipart body
	if err = writer.Close(); err != nil {
		return fmt.Errorf("failed to close multipart writer: %w", err)
	}

	// Create HTTP request with Content-Length
	req, err := http.NewRequest(http.MethodPost, url, &b)
	if err != nil {
		return fmt.Errorf("failed to create HTTP request: %w", err)
	}
	req.Header.Set(HeaderContentType, writer.FormDataContentType())

	// Use an HTTP client with a timeout
	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode >= 300 {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("upload failed with status %d: %s", resp.StatusCode, string(body))
	}

	return nil
}
