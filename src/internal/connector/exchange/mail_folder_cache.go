package exchange

import (
	"context"

	multierror "github.com/hashicorp/go-multierror"
	msfolderdelta "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders/item/childfolders/delta"
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
	cache          map[string]graph.CachedContainer
	gs             graph.Service
	userID, rootID string
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

	// Root only needs the ID because we hide it's name for Mail.
	idPtr := f.GetId()

	if idPtr == nil || len(*idPtr) == 0 {
		return errors.New("root folder has no ID")
	}

	temp := cacheFolder{
		Container: f,
		p:         path.Builder{}.Append(baseContainerPath...),
	}
	mc.cache[*idPtr] = &temp
	mc.rootID = *idPtr

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
		MailFoldersById(mc.rootID).ChildFolders().
		Delta()

	var errs *multierror.Error

	// TODO: Cannot use Iterator for delta
	// Awaiting resolution: https://github.com/microsoftgraph/msgraph-sdk-go/issues/272
	for {
		resp, err := query.Get(ctx, nil)
		if err != nil {
			return errors.Wrap(err, support.ConnectorStackErrorTrace(err))
		}

		for _, f := range resp.GetValue() {
			if err := mc.AddToCache(ctx, f); err != nil {
				errs = multierror.Append(errs, errors.Wrap(err, "error on delta fetch"))
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

	return errs.ErrorOrNil()
}

func (mc *mailFolderCache) IDToPath(
	ctx context.Context,
	folderID string,
) (*path.Builder, error) {
	c, ok := mc.cache[folderID]
	if !ok {
		return nil, errors.Errorf("folder %s not cached", folderID)
	}

	p := c.Path()
	if p != nil {
		return p, nil
	}

	parentPath, err := mc.IDToPath(ctx, *c.GetParentFolderId())
	if err != nil {
		return nil, errors.Wrap(err, "retrieving parent folder")
	}

	fullPath := parentPath.Append(*c.GetDisplayName())
	c.SetPath(fullPath)

	return fullPath, nil
}

// init ensures that the structure's fields are initialized.
// Fields Initialized when cache == nil:
// [mc.cache, mc.rootID]
func (mc *mailFolderCache) init(
	ctx context.Context,
	baseNode string,
	baseContainerPath []string,
) error {
	if len(baseNode) == 0 {
		return errors.New("m365 folder ID required for base folder")
	}

	if mc.cache == nil {
		mc.cache = map[string]graph.CachedContainer{}
	}

	return mc.populateMailRoot(ctx, baseNode, baseContainerPath)
}

// AddToCache adds container to map in field 'cache'
// @returns error iff the required values are not accessible.
func (mc *mailFolderCache) AddToCache(ctx context.Context, f graph.Container) error {
	if err := checkRequiredValues(f); err != nil {
		return errors.Wrap(err, "object not added to cache")
	}

	if _, ok := mc.cache[*f.GetId()]; ok {
		return nil
	}

	mc.cache[*f.GetId()] = &cacheFolder{
		Container: f,
	}

	// Populate the path for this entry so calls to PathInCache succeed no matter
	// when they're made.
	_, err := mc.IDToPath(ctx, *f.GetId())
	if err != nil {
		return errors.Wrap(err, "adding cache entry")
	}

	return nil
}

// PathInCache utility function to return m365ID of folder if the pathString
// matches the path of a container within the cache.
func (mc *mailFolderCache) PathInCache(pathString string) (string, bool) {
	if len(pathString) == 0 || mc.cache == nil {
		return "", false
	}

	for _, folder := range mc.cache {
		if folder.Path() == nil {
			continue
		}

		if folder.Path().String() == pathString {
			return *folder.GetId(), true
		}
	}

	return "", false
}

func (mc *mailFolderCache) Items() []graph.CachedContainer {
	res := make([]graph.CachedContainer, 0, len(mc.cache))

	for _, c := range mc.cache {
		res = append(res, c)
	}

	return res
}
