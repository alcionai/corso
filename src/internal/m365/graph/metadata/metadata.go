package metadata

import (
	"github.com/alcionai/corso/src/internal/m365/collection/drive/metadata"
	"github.com/alcionai/corso/src/pkg/path"
)

// IsMetadataFilePath checks whether the LAST service in the path
// supports metadata file types and, if so, whether the item has
// a meta suffix.
func IsMetadataFilePath(p path.Path) bool {
	return IsMetadataFile(
		p.ServiceResources(),
		p.Category(),
		p.Item())
}

// IsMetadataFile accepts the ServiceResources, cat, and Item values from
// a path (or equivalent representation) and returns true if the item
// is a Metadata entry.
func IsMetadataFile(
	srs []path.ServiceResource,
	cat path.CategoryType,
	itemID string,
) bool {
	switch srs[len(srs)-1].Service {
	case path.OneDriveService:
		return metadata.HasMetaSuffix(itemID)

	case path.SharePointService:
		return cat == path.LibrariesCategory && metadata.HasMetaSuffix(itemID)

	default:
		return false
	}
}
