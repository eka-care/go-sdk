package records

import (
	"bytes"
	"io"
	"net/http"
)

func GetFileSize(file io.Reader) int {
	buf := new(bytes.Buffer)
	buf.ReadFrom(file)
	return buf.Len()
}

func GetContentType(content []byte) string {
	return http.DetectContentType(content)
}
