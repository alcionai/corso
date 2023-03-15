package exchange

import (
	"context"

	"github.com/alcionai/clues"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
)

var _ graph.ContainerResolver = &mailFolderCache{}

// mailFolderCache struct used to improve lookup of directories within exchange.Mail
// cache map of cachedContainers where the  key =  M365ID
// nameLookup map: Key: DisplayName Value: ID
type mailFolderCache struct {
	*containerResolver
	enumer containersEnumerator
	getter containerGetter
	userID string
}

// init ensures that the structure's fields are initialized.
// Fields Initialized when cache == nil:
// [mc.cache]
func (mc *mailFolderCache) init(
	ctx context.Context,
) error {
	if mc.containerResolver == nil {
		mc.containerResolver = newContainerResolver()
	}

	return mc.populateMailRoot(ctx)
}

// populateMailRoot manually fetches directories that are not returned during Graph for msgraph-sdk-go v. 40+
// rootFolderAlias is the top-level directory for exchange.Mail.
// Action ensures that cache will stop at appropriate level.
// @error iff the struct is not properly instantiated
func (mc *mailFolderCache) populateMailRoot(ctx context.Context) error {
	f, err := mc.getter.GetContainerByID(ctx, mc.userID, rootFolderAlias)
	if err != nil {
		return clues.Wrap(err, "fetching root folder")
	}

	temp := graph.NewCacheFolder(
		f,
		// Root folder doesn't store any mail messages so it isn't given any paths.
		// Giving it a path/location would cause us to add extra path elements that
		// the user doesn't see in the regular UI for Exchange.
		path.Builder{}.Append(), // storage path
		path.Builder{}.Append()) // display location
	if err := mc.addFolder(temp); err != nil {
		return clues.Wrap(err, "adding resolver dir").WithClues(ctx)
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
	errs *fault.Bus,
	baseID string,
	baseContainerPath ...string,
) error {
	if err := mc.init(ctx); err != nil {
		return errors.Wrap(err, "initializing")
	}

	err := mc.enumer.EnumerateContainers(ctx, mc.userID, "", mc.addFolder, errs)
	if err != nil {
		return errors.Wrap(err, "enumerating containers")
	}

	if err := mc.populatePaths(ctx, false); err != nil {
		return errors.Wrap(err, "populating paths")
	}

	return nil
}
