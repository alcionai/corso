package site

import (
	"time"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/internal/m365/graph/betasdk/models"
	"github.com/alcionai/corso/src/pkg/backup/details"
)

// pageToSPInfo propagates metadata from  the SharePoint Page data type
// into searchable content.
// Page Details: https://learn.microsoft.com/en-us/graph/api/resources/sitepage?view=graph-rest-beta
func pageToSPInfo(page models.SitePageable, root string, size int64) *details.SharePointInfo {
	var (
		name, prefix, webURL string
		created, modified    time.Time
	)

	if title, ok := ptr.ValOK(page.GetTitle()); ok {
		name = title
	}

	if page.GetWebUrl() != nil {
		if len(root) > 0 {
			prefix = root + "/"
		}

		webURL = prefix + ptr.Val(page.GetWebUrl())
	}

	if page.GetCreatedDateTime() != nil {
		created = ptr.Val(page.GetCreatedDateTime())
	}

	modified = ptr.OrNow(page.GetLastModifiedDateTime())

	return &details.SharePointInfo{
		ItemType:   details.SharePointPage,
		ItemName:   name,
		ParentPath: root,
		Created:    created,
		Modified:   modified,
		WebURL:     webURL,
		Size:       size,
	}
}
