package kopia

import (
	"strings"

	"github.com/alcionai/corso/src/pkg/path"
)

type SubtreeMigrator interface {
	// GetNewSubtree potentially transforms a given subtree (repo tree prefix
	// corresponding to a kopia Reason (eg: resource owner, service, category)
	// into a new subtree when merging items from the base tree.
	GetNewSubtree(oldSubtree *path.Builder) *path.Builder
}

type subtreeOwnerMigrator struct {
	new, old string
}

// migrates any subtree with a matching old owner onto the new owner
func (om subtreeOwnerMigrator) GetNewSubtree(old *path.Builder) *path.Builder {
	if old == nil {
		return nil
	}

	elems := old.Elements()
	if len(elems) < 4 {
		return old
	}

	if strings.EqualFold(elems[2], om.old) {
		elems[2] = om.new
	}

	return path.Builder{}.Append(elems...)
}

func NewSubtreeOwnerMigration(new, old string) *subtreeOwnerMigrator {
	return &subtreeOwnerMigrator{new, old}
}
