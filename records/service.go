package records

import (
	"encoding/json"
	"errors"
	"os"
)

func (r *RecordsService) GetAuthorization(batchRequest []BatchRequest) (*AuthorizationResponse, error) {
	data, _ := json.Marshal(map[string]interface{}{"batch_request": batchRequest})
	resp, err := r.client.Request("POST", r.url, data, nil)
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

func (r *RecordsService) UploadDocument(filePath, documentType string, documentDate *int, tags []string, title *string) (string, error) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return "", errors.New("file not found")
	}

	fileSize, _ := getFileSize(filePath)
	contentType := getContentType(filePath)
	batchRequest := []BatchRequest{{
		DocumentType: documentType,
		DocumentDate: documentDate,
		Tags:         tags,
		Files:        []File{{ContentType: contentType, FileSize: fileSize}},
		Title:        title,
	}}
	authResp, err := r.GetAuthorization(batchRequest)
	if err != nil {
		return "", err
	}

	if len(authResp.BatchResponse) == 0 || len(authResp.BatchResponse[0].Forms) == 0 {
		return "", errors.New("no upload URL received")
	}
	documentID := authResp.BatchResponse[0].DocumentID
	form := authResp.BatchResponse[0].Forms[0]

	return documentID, uploadFile(filePath, form.URL, form.Fields)
}
