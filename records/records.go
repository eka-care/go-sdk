// records.go
package records

import (
	"fmt"

	"github.com/ekacare/go-sdk/client"
)

type RecordsService struct {
	client           *client.Client
	authorizationURL string
}

func NewRecordsService(client *client.Client) *RecordsService {
	authorizationURL := fmt.Sprintf("%s%s", client.BaseURL, UPLOAD_AUTHORIZATION_PATH)
	return &RecordsService{
		client:           client,
		authorizationURL: authorizationURL,
	}
}
