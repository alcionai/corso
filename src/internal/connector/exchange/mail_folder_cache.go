package exchange

import (
	"context"
	"fmt"
	"strings"

	multierror "github.com/hashicorp/go-multierror"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	msfolderdelta "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders/item/childfolders/delta"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/pkg/path"
)

var _ cachedContainer = &mailFolder{}

// cachedContainer is used for local unit tests but also makes it so that this
// code can be broken into generic- and service-specific chunks later on to
// reuse logic in IDToPath.
type cachedContainer interface {
	container
	Path() *path.Builder
	SetPath(*path.Builder)
}

// mailFolder structure that implements the cachedContainer interface
type mailFolder struct {
	folder container
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
	cache          map[string]cachedContainer
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
		return errors.Wrapf(err, "fetching root folder")
	}

	// Root only needs the ID because we hide it's name for Mail.
	idPtr := f.GetId()

	if idPtr == nil || len(*idPtr) == 0 {
		return errors.New("root folder has no ID")
	}

	temp := mailFolder{
		folder: f,
		p:      path.Builder{}.Append(baseContainerPath...),
	}
	mc.cache[*idPtr] = &temp
	mc.rootID = *idPtr

	return nil
}

// checkRequiredValues is a helper function to ensure that
// all the pointers are set prior to being called.
func checkRequiredValues(c container) error {
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
			return err
		}

		for _, f := range resp.GetValue() {
			if err := mc.addMailFolder(f); err != nil {
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

// init ensures that the structure's fields are initialized.
// Fields Initialized when cache == nil:
// [mc.cache, mc.rootID]
func (mc *mailFolderCache) init(
	ctx context.Context,
	baseNode string,
	baseContainerPath []string,
) error {
	if len(baseNode) == 0 {
		return errors.New("M365 folder ID required for base folder")
	}

	if mc.cache == nil {
		mc.cache = map[string]cachedContainer{}
	}

	return mc.populateMailRoot(ctx, baseNode, baseContainerPath)
}

// addMailFolder adds container to map in field 'cache'
// @returns error iff the required values are not accessible.
func (mc *mailFolderCache) addMailFolder(f models.MailFolderable) error {
	if err := checkRequiredValues(f); err != nil {
		return err
	}

	if _, ok := mc.cache[*f.GetId()]; ok {
		return nil
	}

	mc.cache[*f.GetId()] = &mailFolder{
		folder: f,
	}

	return nil
}

// find path
func (mc *mailFolderCache) PathInCache(pathString string) (string, bool) {
	if len(pathString) == 0 || mc.cache == nil {
		return "", false
	}

	for _, folder := range mc.cache {
		if folder.Path() == nil {
			fmt.Println("Exited with an empty path")
			return "", false
		}
		pathRep := folder.Path().String()
		fmt.Printf("Cached: %s\tValue:%s \tMatch: %v\n", pathRep, pathString, pathRep == pathString)
		if pathRep == pathString {
			return *folder.GetId(), true
		}
	}

	return "", false
}

// pathElementStringBuilder helper function for returning
// a string separated with '/' based on the index.
// Returns full slice separated w/ '/' if index is Greather Than or Equal
// to the length of the slice.
func pathElementStringBuilder(index int, slice []string) string {
	if index >= len(slice) {
		return strings.Join(slice, "/")
	}

	return strings.Join(slice[:index], "/")
}
