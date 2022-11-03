package exchange

import (
	"context"

	msgraphgocore "github.com/microsoftgraph/msgraph-sdk-go-core"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/pkg/path"
)

var _ graph.ContainerResolver = &contactFolderCache{}

type contactFolderCache struct {
	*containerResolver
	gs     graph.Service
	userID string
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
	if err != nil {
		return errors.Wrapf(
			err,
			"fetching root contact folder: "+support.ConnectorStackErrorTrace(err))
	}

	temp := cacheFolder{
		Container: f,
		p:         path.Builder{}.Append(baseContainerPath...),
	}

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

	var (
		containers = make(map[string]graph.Container)
		errs       error
		errUpdater = func(s string, e error) {
			errs = support.WrapAndAppend(s, e, errs)
		}
	)

	query, err := cfc.
		gs.Client().
		UsersById(cfc.userID).
		ContactFoldersById(baseID).
		ChildFolders().
		Get(ctx, nil)
	if err != nil {
		return errors.Wrap(err, support.ConnectorStackErrorTrace(err))
	}

	iter, err := msgraphgocore.NewPageIterator(query, cfc.gs.Adapter(),
		models.CreateContactFolderCollectionResponseFromDiscriminatorValue)
	if err != nil {
		return errors.Wrap(err, support.ConnectorStackErrorTrace(err))
	}

	cb := IterativeCollectContactContainers(containers,
		"",
		errUpdater)
	if err := iter.Iterate(ctx, cb); err != nil {
		return errors.Wrap(err, support.ConnectorStackErrorTrace(err))
	}

	if errs != nil {
		return errs
	}

	for _, entry := range containers {
		temp := cacheFolder{
			Container: entry,
		}

		err = cfc.addFolder(temp)
		if err != nil {
			errs = support.WrapAndAppend(
				"cache build in cfc.Populate",
				err,
				errs)
		}
	}

	if err := cfc.populatePaths(ctx); err != nil {
		errs = support.WrapAndAppend(
			"contacts resolver",
			err,
			errs,
		)
	}

	return errs
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
