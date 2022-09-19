package exchange

import (
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/internal/path"
	"github.com/alcionai/corso/src/pkg/backup/details"
)

// ContactInfo translate models.Contactable metadata into searchable content
func ContactInfo(contact models.Contactable) *details.ExchangeInfo {
	name := ""

	if contact.GetDisplayName() != nil {
		name = *contact.GetDisplayName()
	}

	return &details.ExchangeInfo{
		ItemType:    path.ContactsCategory,
		ContactName: name,
	}
}
