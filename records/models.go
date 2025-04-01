package records

import "io"

type DocumentTypeQueryParam string

const PrescriptionQP DocumentTypeQueryParam = "ps"
const LabReportQP DocumentTypeQueryParam = "lr"
const OtherQP DocumentTypeQueryParam = "ot"
const DischargeSummaryQP DocumentTypeQueryParam = "ds"
const VaccineCertificateQP DocumentTypeQueryParam = "vc"
const InsuranceQP DocumentTypeQueryParam = "in"
const InvoiceQP DocumentTypeQueryParam = "iv"
const ScanQP DocumentTypeQueryParam = "sc"

func (d DocumentTypeQueryParam) AsP() *DocumentTypeQueryParam {
	return &d
}

type File struct {
	Content     io.Reader `json:"-"`
	ContentType string    `json:"contentType"`
	FileSize    int64     `json:"file_size"`
}

const SmartReportTaskQP Task = "smart"
const PIITaskQP Task = "pii"
const ClassificationTaskQP Task = "classification"

type Task string

func (t Task) AsP() *Task {
	return &t
}

type BatchRequest struct {
	DocumentType DocumentTypeQueryParam `json:"dt"`
	DocumentDate *int64                 `json:"dd_e,omitempty"`
	Tags         []string               `json:"tg,omitempty"`
	Files        []File                 `json:"files"`
	Title        string                 `json:"title,omitempty"`
}

type UploadRequest struct {
	Request []BatchRequest `json:"request"`
	Tasks   []Task         `json:"-"`
	Batch   bool           `json:"-"`
}

type authorizationResponse struct {
	BatchResponse []struct {
		DocumentID string `json:"document_id"`
		Forms      []struct {
			URL    string            `json:"url"`
			Fields map[string]string `json:"fields"`
		} `json:"forms"`
	} `json:"batch_response"`
}

type UploadResponse struct {
	DocumentIDs []*string `json:"document_ids"`
}
