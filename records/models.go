package records

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
	Content     []byte `json:"-"`
	ContentType string `json:"contentType"`
	FileSize    int    `json:"file_size"`
}

type BatchRequest struct {
	DocumentType DocumentTypeQueryParam `json:"dt"`
	DocumentDate *int64                 `json:"dd_e,omitempty"`
	Tags         []string               `json:"tg,omitempty"`
	Files        []File                 `json:"files"`
	Title        string                 `json:"title,omitempty"`
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
