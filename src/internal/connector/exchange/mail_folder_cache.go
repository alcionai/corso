package exchange

import (
	"context"

	multierror "github.com/hashicorp/go-multierror"
	msfolderdelta "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders/item/childfolders/delta"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/pkg/path"
)

var _ graph.CachedContainer = &mailFolder{}

// mailFolder structure that implements the graph.CachedContainer interface
type mailFolder struct {
	folder graph.Container
	p      *path.Builder
}

//=========================================
// Required Functions to satisfy interfaces
//=====================================

func (mf mailFolder) Path() *path.Builder {
	return mf.p
}

func (mf *mailFolder) SetPath(newPath *path.Builder) {
	mf.p = newPath
}

func (mf *mailFolder) GetDisplayName() *string {
	return mf.folder.GetDisplayName()
}

//nolint:revive
func (mf *mailFolder) GetId() *string {
	return mf.folder.GetId()
}

//nolint:revive
func (mf *mailFolder) GetParentFolderId() *string {
	return mf.folder.GetParentFolderId()
}

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
func (mc *mailFolderCache) populateMailRoot(ctx context.Context, directoryID string) error {
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
		return errors.Wrapf(err, "fetching root folder")
	}

	// Root only needs the ID because we hide it's name for Mail.
	idPtr := f.GetId()

	if idPtr == nil || len(*idPtr) == 0 {
		return errors.New("root folder has no ID")
	}

	temp := mailFolder{
		folder: f,
		p:      &path.Builder{},
	}
	mc.cache[*idPtr] = &temp
	mc.rootID = *idPtr

	return nil
}

// checkRequiredValues is a helper function to ensure that
// all the pointers are set prior to being called.
func checkRequiredValues(c graph.Container) error {
	idPtr := c.GetId()
	if idPtr == nil || len(*idPtr) == 0 {
		return errors.New("folder without ID")
	}

	ptr := c.GetDisplayName()
	if ptr == nil || len(*ptr) == 0 {
		return errors.Errorf("folder %s without display name", *idPtr)
	}

	ptr = c.GetParentFolderId()
	if ptr == nil || len(*ptr) == 0 {
		return errors.Errorf("folder %s without parent ID", *idPtr)
	}

	return nil
}

// Populate utility function for populating the mailFolderCache.
// Number of Graph Queries: 1.
// @param baseID: M365ID of the base of the exchange.Mail.Folder
// Use rootFolderAlias for input if baseID unknown
func (mc *mailFolderCache) Populate(ctx context.Context, baseID string) error {
	if len(baseID) == 0 {
		return errors.New("populate function requires: M365ID as input")
	}

	err := mc.Init(ctx, baseID)
	if err != nil {
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
			return err
		}

		for _, f := range resp.GetValue() {
			if err := mc.AddToCache(ctx, f); err != nil {
				errs = multierror.Append(errs, err)
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

// Init ensures that the structure's fields are initialized.
// Fields Initialized when cache == nil:
// [mc.cache, mc.rootID]
func (mc *mailFolderCache) Init(ctx context.Context, baseNode string) error {
	if mc.cache == nil {
		mc.cache = map[string]graph.CachedContainer{}
	}

	return mc.populateMailRoot(ctx, baseNode)
}

func (mc *mailFolderCache) AddToCache(ctx context.Context, f graph.Container) error {
	if err := checkRequiredValues(f); err != nil {
		return errors.Wrap(err, "adding cache entry")
	}

	if _, ok := mc.cache[*f.GetId()]; ok {
		return nil
	}

	mc.cache[*f.GetId()] = &mailFolder{
		folder: f,
	}

	_, err := mc.IDToPath(ctx, *f.GetId())
	if err != nil {
		return errors.Wrap(err, "adding cache entry")
	}

	return nil
}

func (mc *mailFolderCache) Items() []graph.CachedContainer {
	res := make([]graph.CachedContainer, 0, len(mc.cache))

	for _, c := range mc.cache {
		res = append(res, c)
	}

	return res
}
