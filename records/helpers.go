package records

import (
	"bytes"
	"io"
	"path/filepath"
)

func GetFileSize(file io.Reader) int {
	buf := new(bytes.Buffer)
	buf.ReadFrom(file)
	return buf.Len()
}

func GetContentType(filePath string) string {
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
