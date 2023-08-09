package metadata

import (
	"github.com/alcionai/corso/src/internal/m365/service/onedrive/metadata"
	"github.com/alcionai/corso/src/pkg/path"
)

func IsMetadataFile(p path.Path) bool {
	switch p.Service() {
	case path.OneDriveService:
		return metadata.HasMetaSuffix(p.Item())

	case path.SharePointService:
		return p.Category() == path.LibrariesCategory && metadata.HasMetaSuffix(p.Item())

	default:
		return false
	}
}
