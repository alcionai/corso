package drive

import (
	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/pkg/path"
)

// TODO: refactor to comply with graph/cache_container

type folderCache struct {
	cache map[string]models.DriveItemable
}

func NewFolderCache() *folderCache {
	return &folderCache{
		cache: map[string]models.DriveItemable{},
	}
}

func (c *folderCache) get(loc *path.Builder) (models.DriveItemable, bool) {
	mdi, ok := c.cache[loc.String()]
	return mdi, ok
}

func (c *folderCache) set(loc *path.Builder, mdi models.DriveItemable) {
	c.cache[loc.String()] = mdi
}
