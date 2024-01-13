package metadata

import (
	"github.com/alcionai/corso/src/internal/m365/collection/drive/metadata"
	"github.com/alcionai/corso/src/pkg/path"
)

func IsMetadataFile(p path.Path) bool {
	switch p.Service() {
	case path.OneDriveService:
		return metadata.HasMetaSuffix(p.Item())

	case path.SharePointService, path.GroupsService:
		return p.Category() == path.LibrariesCategory && metadata.HasMetaSuffix(p.Item())

	default:
		return false
	}
}
