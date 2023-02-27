package metadata

import (
	"github.com/alcionai/corso/src/internal/connector/onedrive"
	"github.com/alcionai/corso/src/pkg/path"
)

func IsMetadataFile(p path.Path) bool {
	switch p.Service() {
	case path.OneDriveService:
		return onedrive.IsMetaFile(p.Item())

	default:
		return false
	}
}
