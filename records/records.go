// records.go
package records

import (
	"fmt"

	"github.com/ekacare/go-sdk/client"
)

type RecordsService struct {
	client *client.Client
	url    string
}

func NewRecordsService(client *client.Client) *RecordsService {
	url := fmt.Sprintf("%s%s", client.BaseURL, UPLOAD_AUTHORIZATION_PATH)
	return &RecordsService{
		client: client,
		url:    url,
	}
}
