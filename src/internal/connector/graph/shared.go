package graph

// Contains helper methods that are common across multiple
// Microsoft Graph Applications.

import (
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
)

func UnwrapEmailAddress(contact models.Recipientable) string {
	var empty string
	if contact == nil || contact.GetEmailAddress() == nil {
		return empty
	}

	return ptr.Val(contact.GetEmailAddress().GetAddress())
}
