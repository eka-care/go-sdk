package records

import "io"

type BatchRequest struct {
	DocumentType string   `json:"dt"`
	DocumentDate *int     `json:"dd_e,omitempty"`
	Tags         []string `json:"tg,omitempty"`
	Files        []File   `json:"files"`
	Title        *string  `json:"title,omitempty"`
}

type File struct {
	ContentType string `json:"contentType"`
	FileSize    int    `json:"file_size"`
}

type FileRequest struct {
	FileName string
	Reader   io.Reader
}

type AuthorizationResponse struct {
	BatchResponse []struct {
		DocumentID string `json:"document_id"`
		Forms      []struct {
			URL    string            `json:"url"`
			Fields map[string]string `json:"fields"`
		} `json:"forms"`
	} `json:"batch_response"`
}
