// records.go
package records

import (
	"github.com/ekacare/go-sdk/client"
)

type Records interface {
	UploadDocument(batchRequest UploadRequest) (*UploadResponse, error)
}

type RecordsService struct {
	client client.ClientInterface
}

func NewRecordsService(client client.ClientInterface) Records {
	return &RecordsService{
		client: client,
	}
}
