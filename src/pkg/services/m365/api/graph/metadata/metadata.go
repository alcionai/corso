package metadata

import (
	"github.com/alcionai/corso/src/pkg/path"
)

func IsMetadataFile(p path.Path) bool {
	switch p.Service() {
	case path.OneDriveService:
		return HasMetaSuffix(p.Item())

	case path.SharePointService, path.GroupsService:
		return p.Category() == path.LibrariesCategory && HasMetaSuffix(p.Item())

	default:
		return false
	}
}
