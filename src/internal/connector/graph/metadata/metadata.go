package metadata

import (
	"github.com/alcionai/corso/src/internal/connector/onedrive/common"
	"github.com/alcionai/corso/src/pkg/path"
)

func IsMetadataFile(p path.Path) bool {
	switch p.Service() {
	case path.OneDriveService:
		return common.IsMetaFile(p.Item())

	default:
		return false
	}
}
