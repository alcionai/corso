package exchange

import (
	"context"

	multierror "github.com/hashicorp/go-multierror"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
	msfolderdelta "github.com/microsoftgraph/msgraph-sdk-go/users/item/mailfolders/delta"
	"github.com/pkg/errors"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/pkg/path"
)

// cachedContainer is used for local unit tests but also makes it so that this
// code can be broken into generic- and service-specific chunks later on to
// reuse logic in IDToPath.
type cachedContainer interface {
	container
	Path() *path.Builder
	SetPath(*path.Builder)
}

type mailFolder struct {
	models.MailFolderable
	p *path.Builder
}

func (mf mailFolder) Path() *path.Builder {
	return mf.p
}

func (mf *mailFolder) SetPath(newPath *path.Builder) {
	mf.p = newPath
}

type mailFolderCache struct {
	cache  map[string]cachedContainer
	gs     graph.Service
	userID string
}

// populateRoot fetches and populates the root folder in the cache so the cache
// knows when to stop resolving the path.
func (mc *mailFolderCache) populateRoot(ctx context.Context) error {
	wantedOpts := []string{"displayName", "parentFolderId"}

	opts, err := optionsForMailFoldersItem(wantedOpts)
	if err != nil {
		return errors.Wrapf(err, "getting options for mail folders %v", wantedOpts)
	}

	f, err := mc.
		gs.
		Client().
		UsersById(mc.userID).
		MailFoldersById(rootFolderAlias).
		Get(ctx, opts)
	if err != nil {
		return errors.Wrapf(err, "fetching root folder")
	}

	// Root only needs the ID because we hide it's name for Mail.
	idPtr := f.GetId()
	if idPtr == nil || len(*idPtr) == 0 {
		return errors.New("root folder has no ID")
	}

	mc.cache[*idPtr] = &mailFolder{
		MailFolderable: f,
		p:              &path.Builder{},
	}

	return nil
}

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

func (mc *mailFolderCache) Populate(ctx context.Context) error {
	if mc.cache == nil {
		mc.cache = map[string]cachedContainer{}
	}

	if err := mc.populateRoot(ctx); err != nil {
		return err
	}

	builder := mc.
		gs.
		Client().
		UsersById(mc.userID).
		MailFolders().
		Delta()

	var errs *multierror.Error

	for {
		resp, err := builder.Get(ctx, nil)
		if err != nil {
			return err
		}

		for _, f := range resp.GetValue() {
			if err := checkRequiredValues(f); err != nil {
				errs = multierror.Append(errs, err)
				continue
			}

			mc.cache[*f.GetId()] = &mailFolder{MailFolderable: f}
		}

		r := resp.GetAdditionalData()

		n, ok := r[nextDataLink]
		if !ok || n == nil {
			break
		}

		link := *(n.(*string))
		builder = msfolderdelta.NewDeltaRequestBuilder(link, mc.gs.Adapter())
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
