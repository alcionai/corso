package graph

import (
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/pkg/path"
)

// CachedContainer is used for local unit tests but also makes it so that this
// code can be broken into generic- and service-specific chunks later on to
// reuse logic in IDToPath.
type CachedContainer interface {
	Container
	Path() *path.Builder
	SetPath(*path.Builder)
}

// CheckIDAndName is a helper function to ensure that
// the ID and name pointers are set prior to being called.
func CheckIDAndName(c Container) error {
	idPtr := c.GetId()
	if idPtr == nil || len(*idPtr) == 0 {
		return errors.New("folder without ID")
	}

	ptr := c.GetDisplayName()
	if ptr == nil || len(*ptr) == 0 {
		return errors.Errorf("folder %s without display name", *idPtr)
	}

	return nil
}

// CheckRequiredValues is a helper function to ensure that
// all the pointers are set prior to being called.
func CheckRequiredValues(c Container) error {
	if err := CheckIDAndName(c); err != nil {
		return err
	}

	ptr := c.GetParentFolderId()
	if ptr == nil || len(*ptr) == 0 {
		return errors.Errorf("folder %s without parent ID", *c.GetId())
	}

	return nil
}

//======================================
// cachedContainer Implementations
//======================================

var _ CachedContainer = &CacheFolder{}

type CacheFolder struct {
	Container
	p *path.Builder
}

// NewCacheFolder public constructor for struct
func NewCacheFolder(c Container, pb *path.Builder) CacheFolder {
	cf := CacheFolder{
		Container: c,
		p:         pb,
	}

	return cf
}

func (cf CacheFolder) Path() *path.Builder {
	return cf.p
}

func (cf *CacheFolder) SetPath(newPath *path.Builder) {
	cf.p = newPath
}
