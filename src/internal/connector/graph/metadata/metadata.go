package metadata

import (
	"github.com/alcionai/corso/src/internal/connector/onedrive/metadata"
	"github.com/alcionai/corso/src/pkg/path"
)

func IsMetadataFile(p path.Path) bool {
	switch p.Service() {
	case path.OneDriveService, path.SharePointService:
		return metadata.HasMetaSuffix(p.Item())

	default:
		return false
	}
}
