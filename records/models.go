package records

import "io"

type BatchRequest struct {
	DocumentType DocumentTypeQueryParam `json:"dt"`
	DocumentDate *int                   `json:"dd_e,omitempty"`
	Tags         []string               `json:"tg,omitempty"`
	Files        []File                 `json:"files"`
	Title        *string                `json:"title,omitempty"`
}

type File struct {
	ContentType string `json:"contentType"`
	FileSize    int    `json:"file_size"`
}

type FileRequest struct {
	FileName     string
	Reader       io.Reader
	DocumentType DocumentTypeQueryParam
	DocumentDate *int
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

const PrescriptionQP DocumentTypeQueryParam = "ps"
const LabReportQP DocumentTypeQueryParam = "lr"
const CowinCertificateQP DocumentTypeQueryParam = "cc"
const CowinAppointmentSlipQP DocumentTypeQueryParam = "ca"
const ProfilePicQP DocumentTypeQueryParam = "pp"
const OtherQP DocumentTypeQueryParam = "ot"
const DischargeSummaryQP DocumentTypeQueryParam = "ds"
const VaccineCertificateQP DocumentTypeQueryParam = "vc"
const InsuranceQP DocumentTypeQueryParam = "in"
const InvoiceQP DocumentTypeQueryParam = "iv"
const ScanQP DocumentTypeQueryParam = "sc"
const NDHMIDCardQP DocumentTypeQueryParam = "nc"
const NDHMQRCodeQP DocumentTypeQueryParam = "nq"

type DocumentTypeQueryParam string

func (d DocumentTypeQueryParam) AsP() *DocumentTypeQueryParam {
	return &d
}
