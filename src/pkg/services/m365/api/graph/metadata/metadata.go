package metadata

import (
	"github.com/alcionai/canario/src/pkg/path"
)

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
