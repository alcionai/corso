package sharepoint

import (
	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/connector/graph/betasdk/models"
	"github.com/alcionai/corso/src/pkg/backup/details"
)

// sharePointPageInfo propagates metadata from  the SharePoint Page data type
// into searchable content.
// Page Details: https://learn.microsoft.com/en-us/graph/api/resources/sitepage?view=graph-rest-beta
func sharePointPageInfo(page models.SitePageable, size int64) *details.SharePointInfo {
	var (
		name     = ptr.Val(page.GetTitle())
		webURL   = ptr.Val(page.GetWebUrl())
		created  = ptr.Val(page.GetCreatedDateTime())
		modified = ptr.Val(page.GetLastModifiedDateTime())
	)

	return &details.SharePointInfo{
		ItemType: details.SharePointItem,
		ItemName: name,
		Created:  created,
		Modified: modified,
		WebURL:   webURL,
		Size:     size,
	}
}
