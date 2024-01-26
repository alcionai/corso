package metadata

import (
	"github.com/alcionai/clues"

	"github.com/alcionai/corso/src/pkg/path"
)

var ErrMetadataFilesNotSupported = clues.New("metadata files not supported")

func IsMetadataFile(p path.Path) bool {
	switch p.Service() {
	case path.OneDriveService:
		return HasMetaSuffix(p.Item())

	case path.SharePointService:
		return p.Category() == path.LibrariesCategory && HasMetaSuffix(p.Item())

	case path.GroupsService:
		return p.Category() == path.LibrariesCategory && HasMetaSuffix(p.Item()) ||
			p.Category() == path.ConversationPostsCategory && HasMetaSuffix(p.Item())
	default:
		return false
	}
}

// func WithDataSuffix(p path.Path) path.Path {
// 	return p.WithItem(p.Item() + ".data")
// }
