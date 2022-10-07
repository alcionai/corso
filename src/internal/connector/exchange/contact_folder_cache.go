package exchange

import (
	"context"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/pkg/errors"
)

var _ cachedContainer = &contactFolder{}

type contactFolder struct {
	container
	p *path.Builder
}

func (cf contactFolder) Path() *path.Builder {
	return cf.p
}

func (cf *contactFolder) SetPath(newPath *path.Builder) {
	cf.p = newPath
}

type contactFolderCache struct {
	cache          map[string]cachedContainer
	gs             graph.Service
	userID, rootID string
}

func (cfc *contactFolderCache) populateContactRoot(
	ctx context.Context,
	directoryID string,
	baseContainerPath []string,
) error {
	wantedOpts := []string{"displayName", "parentFolderId"}

	opts, err := optionsForContactFolderByID(wantedOpts)
	if err != nil {
		return errors.Wrapf(err, "getting options for contact folder cache: %v", wantedOpts)
	}

	f, err := cfc.
		gs.
		Client().
		UsersById(cfc.userID).
		ContactFoldersById(directoryID).
		Get(ctx, opts)
}
