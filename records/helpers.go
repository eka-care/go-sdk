package records

import (
	"bytes"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

func getFileSize(filePath string) (int, error) {
	file, err := os.Stat(filePath)
	if err != nil {
		return 0, err
	}

	return int(file.Size()), nil
}

func getContentType(filePath string) string {
	ext := filepath.Ext(filePath)
	switch ext {
	case ".jpg", ".jpeg":
		return "image/jpeg"
	case ".png":
		return "image/png"
	case ".pdf":
		return "application/pdf"
	default:
		return "application/octet-stream"
	}
}

func uploadFile(filePath, url string, fields map[string]string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	var b bytes.Buffer
	writer := multipart.NewWriter(&b)

	for key, val := range fields {
		_ = writer.WriteField(key, val)
	}

	part, _ := writer.CreateFormFile("file", filepath.Base(filePath))
	_, _ = io.Copy(part, file)
	writer.Close()

	req, _ := http.NewRequest("POST", url, &b)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 300 {
		return errors.New("upload failed")
	}
	return nil
}
