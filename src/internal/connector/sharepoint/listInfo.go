package sharepoint

import (
	"time"

	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/pkg/backup/details"
)

// sharePointListInfo translates models.Listable metadata into searchable content
// List Details: https://learn.microsoft.com/en-us/graph/api/resources/list?view=graph-rest-1.0
func sharePointListInfo(lst models.Listable, size int64) *details.SharePointInfo {
	var (
		name, webURL      string
		created, modified time.Time
	)

	if lst.GetDisplayName() != nil {
		name = *lst.GetDisplayName()
	}

	if lst.GetWebUrl() != nil {
		webURL = *lst.GetWebUrl()
	}

	if lst.GetCreatedDateTime() != nil {
		created = *lst.GetCreatedDateTime()
	}

	if lst.GetLastModifiedDateTime() != nil {
		modified = *lst.GetLastModifiedDateTime()
	}

	return &details.SharePointInfo{
		ItemType: details.SharePointItem,
		ItemName: name,
		Created:  created,
		Modified: modified,
		WebURL:   webURL,
		Size:     size,
	}
}
