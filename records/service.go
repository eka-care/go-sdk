package records

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"time"
)

func (r *RecordsService) GetAuthorization(batchRequest []BatchRequest) (*AuthorizationResponse, error) {
	data, _ := json.Marshal(map[string]interface{}{"batch_request": batchRequest})
	resp, err := r.client.Request("POST", r.authorizationURL, data, nil)
	if err != nil {
		return nil, err
	}

	var authResponse AuthorizationResponse
	err = json.NewDecoder(resp.Body).Decode(&authResponse)
	if err != nil {
		return nil, err
	}
	return &authResponse, nil
}

func (r *RecordsService) UploadDocument(files []FileRequest, batchRequest []BatchRequest) ([]string, error) {
	authResp, err := r.GetAuthorization(batchRequest)
	if err != nil {
		return nil, err
	}

	if len(authResp.BatchResponse) == 0 {
		return nil, errors.New("no upload URL received")
	}

	var documentIDs []string
	for index, batch := range authResp.BatchResponse {
		if len(batch.Forms) == 0 {
			continue
		}

		documentID := batch.DocumentID
		form := batch.Forms[0]
		err = r.upload(files[index].Content, files[index].FileName, form.URL, form.Fields)

		if err != nil {
			return nil, err
		}

		documentIDs = append(documentIDs, documentID)
	}

	return documentIDs, nil
}

func (r *RecordsService) upload(file []byte, fileName, url string, fields map[string]string) error {
	var b bytes.Buffer
	writer := multipart.NewWriter(&b)

	// Write form fields
	for key, val := range fields {
		if err := writer.WriteField(key, val); err != nil {
			return err
		}
	}

	// Create form file field
	part, err := writer.CreateFormFile("file", fileName)
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
	req, err := http.NewRequest("POST", url, &b)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.ContentLength = int64(b.Len()) // Set Content-Length explicitly

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
