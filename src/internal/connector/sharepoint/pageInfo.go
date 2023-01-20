package sharepoint

import (
	"time"

	"github.com/microsoftgraph/msgraph-beta-sdk-go/models"

	"github.com/alcionai/corso/src/pkg/backup/details"
)

// sharePointPageInfo propagates metadata from  the SharePoint Page data type
// into searchable content.
// Page Details: https://learn.microsoft.com/en-us/graph/api/resources/sitepage?view=graph-rest-beta
func sharePointPageInfo(page models.SitePageable, size int64) *details.SharePointInfo {
	var (
		name, webURL      string
		created, modified time.Time
	)

	if page.GetTitle() != nil {
		name = *page.GetTitle()
	}

	if page.GetWebUrl() != nil {
		webURL = *page.GetWebUrl()
	}

	if page.GetCreatedDateTime() != nil {
		created = *page.GetCreatedDateTime()
	}

	if page.GetLastModifiedDateTime() != nil {
		modified = *page.GetLastModifiedDateTime()
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
