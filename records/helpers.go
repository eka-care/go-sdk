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

func GetContentType(reader io.Reader) string {
	// Create a buffer to read the first 512 bytes
	buffer := make([]byte, 512)
	n, err := reader.Read(buffer)
	if err != nil && err != io.EOF {
		return ""
	}

	// Reset the reader to allow re-reading the file
	if seeker, ok := reader.(io.Seeker); ok {
		_, err = seeker.Seek(0, io.SeekStart)
		if err != nil {
			return ""
		}
	} else {
		return ""
	}

	// Detect content type
	return http.DetectContentType(buffer[:n])
}
