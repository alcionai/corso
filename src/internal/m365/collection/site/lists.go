package site

import (
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/internal/common/ptr"
	"github.com/alcionai/corso/src/pkg/backup/details"
)

// ListToSPInfo translates models.Listable metadata into searchable content
// List Details: https://learn.microsoft.com/en-us/graph/api/resources/list?view=graph-rest-1.0
func ListToSPInfo(lst models.Listable, size int64) *details.SharePointInfo {
	var (
		name     = ptr.Val(lst.GetDisplayName())
		webURL   = ptr.Val(lst.GetWebUrl())
		created  = ptr.Val(lst.GetCreatedDateTime())
		modified = ptr.Val(lst.GetLastModifiedDateTime())
	)

	return &details.SharePointInfo{
		ItemType: details.SharePointList,
		ItemName: name,
		Created:  created,
		Modified: modified,
		WebURL:   webURL,
		Size:     size,
	}
}
