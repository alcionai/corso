package exchange

import (
	"time"

	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/pkg/backup/details"
)

// ContactInfo translate models.Contactable metadata into searchable content
func ContactInfo(contact models.Contactable, size int64) *details.ExchangeInfo {
	name := ""
	created := time.Time{}
	modified := time.Time{}

	if contact.GetDisplayName() != nil {
		name = *contact.GetDisplayName()
	}

	if contact.GetCreatedDateTime() != nil {
		created = *contact.GetCreatedDateTime()
	}

	if contact.GetLastModifiedDateTime() != nil {
		modified = *contact.GetLastModifiedDateTime()
	}

	return &details.ExchangeInfo{
		ItemType:    details.ExchangeContact,
		ContactName: name,
		Created:     created,
		Modified:    modified,
		Size:        size,
	}
}
