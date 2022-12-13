package exchange

import (
	"context"

	msuser "github.com/microsoftgraph/msgraph-sdk-go/users"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/pkg/path"
)

var _ graph.ContainerResolver = &contactFolderCache{}

type contactFolderCache struct {
	*containerResolver
	gs     graph.Servicer
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
		errs         error
		options, err = optionsForContactChildFolders([]string{"displayName", "parentFolderId"})
	)

	if err != nil {
		return errors.Wrap(err, "contact cache resolver option")
	}

	builder := cfc.
		gs.Client().
		UsersById(cfc.userID).
		ContactFoldersById(baseID).
		ChildFolders()

	for {
		resp, err := builder.Get(ctx, options)
		if err != nil {
			return errors.Wrap(err, support.ConnectorStackErrorTrace(err))
		}

		for _, fold := range resp.GetValue() {
			if err := checkIDAndName(fold); err != nil {
				errs = support.WrapAndAppend(
					"adding folder to contact resolver",
					err,
					errs,
				)

				continue
			}

			temp := cacheFolder{
				Container: fold,
			}

			err = cfc.addFolder(temp)
			if err != nil {
				errs = support.WrapAndAppend(
					"cache build in cfc.Populate",
					err,
					errs)
			}
		}

		if resp.GetOdataNextLink() == nil {
			break
		}

		builder = msuser.NewUsersItemContactFoldersItemChildFoldersRequestBuilder(*resp.GetOdataNextLink(), cfc.gs.Adapter())
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
