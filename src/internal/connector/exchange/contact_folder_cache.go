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
	cache          map[string]graph.CachedContainer
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
	if err != nil {
		return errors.Wrapf(err, "fetching root contact folder")
	}

	idPtr := f.GetId()

	if idPtr == nil || len(*idPtr) == 0 {
		return errors.New("root folder has no ID")
	}

	temp := cacheFolder{
		Container: f,
		p:         path.Builder{}.Append(baseContainerPath...),
	}
	cfc.cache[*idPtr] = &temp
	cfc.rootID = *idPtr

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
		ContactFoldersById(cfc.rootID).
		ChildFolders().
		Get(ctx, nil)
	if err != nil {
		return errors.Wrap(err, support.ConnectorStackErrorTrace(err))
	}

	iter, err := msgraphgocore.NewPageIterator(query, cfc.gs.Adapter(),
		models.CreateContactFolderCollectionResponseFromDiscriminatorValue)
	if err != nil {
		return err
	}

	cb := IterativeCollectContactContainers(containers,
		"",
		errUpdater)
	if err := iter.Iterate(ctx, cb); err != nil {
		return err
	}

	if errs != nil {
		return errs
	}

	for _, entry := range containers {
		err = cfc.AddToCache(ctx, entry)
		if err != nil {
			errs = support.WrapAndAppend(
				"cache build in cfc.Populate",
				err,
				errs)
		}
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

	if cfc.cache == nil {
		cfc.cache = map[string]graph.CachedContainer{}
	}

	return cfc.populateContactRoot(ctx, baseNode, baseContainerPath)
}

func (cfc *contactFolderCache) IDToPath(
	ctx context.Context,
	folderID string,
) (*path.Builder, error) {
	c, ok := cfc.cache[folderID]
	if !ok {
		return nil, errors.Errorf("folder %s not cached", folderID)
	}

	p := c.Path()
	if p != nil {
		return p, nil
	}

	parentPath, err := cfc.IDToPath(ctx, *c.GetParentFolderId())
	if err != nil {
		return nil, errors.Wrap(err, "retrieving parent folder")
	}

	fullPath := parentPath.Append(*c.GetDisplayName())
	c.SetPath(fullPath)

	return fullPath, nil
}

// PathInCache utility function to return m365ID of folder if the pathString
// matches the path of a container within the cache. A boolean function
// accompanies the call to indicate whether the lookup was successful.
func (cfc *contactFolderCache) PathInCache(pathString string) (string, bool) {
	if len(pathString) == 0 || cfc.cache == nil {
		return "", false
	}

	for _, contain := range cfc.cache {
		if contain.Path() == nil {
			continue
		}

		if contain.Path().String() == pathString {
			return *contain.GetId(), true
		}
	}

	return "", false
}

// AddToCache places container into internal cache field.
// @returns error iff input does not possess accessible values.
func (cfc *contactFolderCache) AddToCache(ctx context.Context, f graph.Container) error {
	if err := checkRequiredValues(f); err != nil {
		return err
	}

	if _, ok := cfc.cache[*f.GetId()]; ok {
		return nil
	}

	cfc.cache[*f.GetId()] = &cacheFolder{
		Container: f,
	}

	// Populate the path for this entry so calls to PathInCache succeed no matter
	// when they're made.
	_, err := cfc.IDToPath(ctx, *f.GetId())
	if err != nil {
		return errors.Wrap(err, "adding cache entry")
	}

	return nil
}

func (cfc *contactFolderCache) Items() []graph.CachedContainer {
	res := make([]graph.CachedContainer, 0, len(cfc.cache))

	for _, c := range cfc.cache {
		res = append(res, c)
	}

	return res
}
