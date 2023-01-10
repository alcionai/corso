package exchange

import (
	"context"

	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/pkg/path"
)

var _ graph.ContainerResolver = &contactFolderCache{}

type contactFolderCache struct {
	*containerResolver
	enumer containersEnumerator
	getter containerGetter
	userID string
}

func (cfc *contactFolderCache) populateContactRoot(
	ctx context.Context,
	directoryID string,
	baseContainerPath []string,
) error {
	f, err := cfc.getter.GetContainerByID(ctx, cfc.userID, directoryID)
	if err != nil {
		return errors.Wrapf(
			err,
			"fetching root contact folder: "+support.ConnectorStackErrorTrace(err))
	}

	temp := graph.NewCacheFolder(f, path.Builder{}.Append(baseContainerPath...))

	if err := cfc.addFolder(temp); err != nil {
		return errors.Wrap(err, "adding cache root")
	}

	return nil
}

// Populate is utility function for placing cache container
// objects into the Contact Folder Cache
// Function does NOT use Delta Queries as it is not supported
// as of (Oct-07-2022)
func (cfc *contactFolderCache) Populate(
	ctx context.Context,
	baseID string,
	baseContainerPather ...string,
) error {
	if err := cfc.init(ctx, baseID, baseContainerPather); err != nil {
		return err
	}

	err := cfc.enumer.EnumerateContainers(ctx, cfc.userID, baseID, cfc.addFolder)
	if err != nil {
		return err
	}

	if err := cfc.populatePaths(ctx); err != nil {
		return errors.Wrap(err, "contacts resolver")
	}

	return nil
}

func (cfc *contactFolderCache) init(
	ctx context.Context,
	baseNode string,
	baseContainerPath []string,
) error {
	if len(baseNode) == 0 {
		return errors.New("m365 folderID required for base folder")
	}

	if cfc.containerResolver == nil {
		cfc.containerResolver = newContainerResolver()
	}

	return cfc.populateContactRoot(ctx, baseNode, baseContainerPath)
}
