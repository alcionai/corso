package exchange

import (
	"context"

	multierror "github.com/hashicorp/go-multierror"
	msfolderdelta "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders/delta"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/internal/connector/support"
	"github.com/alcionai/corso/src/pkg/path"
)

var _ graph.ContainerResolver = &mailFolderCache{}

// mailFolderCache struct used to improve lookup of directories within exchange.Mail
// cache map of cachedContainers where the  key =  M365ID
// nameLookup map: Key: DisplayName Value: ID
type mailFolderCache struct {
	*containerResolver
	gs     graph.Service
	userID string
}

// populateMailRoot fetches and populates the "base" directory from user's inbox.
// Action ensures that cache will stop at appropriate level.
// @param directory: M365 ID of the root all intended inquiries.
// Function should only be used directly when it is known that all
// folder inquiries are going to a specific node. In all other cases
// @error iff the struct is not properly instantiated
func (mc *mailFolderCache) populateMailRoot(
	ctx context.Context,
	directoryID string,
	baseContainerPath []string,
) error {
	wantedOpts := []string{"displayName", "parentFolderId"}

	opts, err := optionsForMailFoldersItem(wantedOpts)
	if err != nil {
		return errors.Wrapf(err, "getting options for mail folders %v", wantedOpts)
	}

	f, err := mc.
		gs.
		Client().
		UsersById(mc.userID).
		MailFoldersById(directoryID).
		Get(ctx, opts)
	if err != nil {
		return errors.Wrap(err, "fetching root folder"+support.ConnectorStackErrorTrace(err))
	}

	temp := cacheFolder{
		Container: f,
		p:         path.Builder{}.Append(baseContainerPath...),
	}

	if err := mc.addFolder(temp); err != nil {
		return errors.Wrap(err, "initializing mail resolver")
	}

	return nil
}

// Populate utility function for populating the mailFolderCache.
// Number of Graph Queries: 1.
// @param baseID: M365ID of the base of the exchange.Mail.Folder
// @param baseContainerPath: the set of folder elements that make up the path
// for the base container in the cache.
func (mc *mailFolderCache) Populate(
	ctx context.Context,
	baseID string,
	baseContainerPath ...string,
) error {
	if err := mc.init(ctx, baseID, baseContainerPath); err != nil {
		return err
	}

	query := mc.
		gs.
		Client().
		UsersById(mc.userID).
		MailFolders().
		Delta()

	var errs *multierror.Error

	for {
		resp, err := query.Get(ctx, nil)
		if err != nil {
			return errors.Wrap(err, support.ConnectorStackErrorTrace(err))
		}

		for _, f := range resp.GetValue() {
			temp := cacheFolder{
				Container: f,
			}

			// Use addFolder instead of AddToCache to be conservative about path
			// population. The fetch order of the folders could cause failures while
			// trying to resolve paths, so put it off until we've gotten all folders.
			if err := mc.addFolder(temp); err != nil {
				errs = multierror.Append(errs, errors.Wrap(err, "delta fetch"))
				continue
			}
		}

		r := resp.GetAdditionalData()

		n, ok := r[nextDataLink]
		if !ok || n == nil {
			break
		}

		link := *(n.(*string))
		query = msfolderdelta.NewDeltaRequestBuilder(link, mc.gs.Adapter())
	}

	if err := mc.populatePaths(ctx); err != nil {
		errs = multierror.Append(errs, errors.Wrap(err, "mail resolver"))
	}

	return errs.ErrorOrNil()
}

// init ensures that the structure's fields are initialized.
// Fields Initialized when cache == nil:
// [mc.cache]
func (mc *mailFolderCache) init(
	ctx context.Context,
	baseNode string,
	baseContainerPath []string,
) error {
	if len(baseNode) == 0 {
		return errors.New("m365 folder ID required for base folder")
	}

	if mc.containerResolver == nil {
		mc.containerResolver = newContainerResolver()
	}

	return mc.populateMailRoot(ctx, baseNode, baseContainerPath)
}
