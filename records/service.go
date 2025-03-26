package records

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
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
		err = r.upload(files[index].Reader, files[index].FileName, form.URL, form.Fields)

		if err != nil {
			return nil, err
		}

		documentIDs = append(documentIDs, documentID)
	}

	return documentIDs, nil
}

func (r *RecordsService) upload(file io.Reader, fileName, url string, fields map[string]string) error {
	var b bytes.Buffer
	writer := multipart.NewWriter(&b)

	for key, val := range fields {
		_ = writer.WriteField(key, val)
	}

	part, _ := writer.CreateFormFile("file", fileName)
	_, _ = io.Copy(part, file)
	writer.Close()

	req, _ := http.NewRequest("POST", url, &b)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 204 {
		return errors.New("upload failed")
	}
	return nil
}
