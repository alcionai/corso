package site

import (
	"context"

	"github.com/microsoftgraph/msgraph-sdk-go/models"

	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

type backupHandler interface {
	getItemByIDer
	getItemser
	canonicalPather
}

// canonicalPath constructs the service and category specific path for
// the given builder.
type canonicalPather interface {
	CanonicalPath(
		storageDir path.Elements,
		tenantID string,
	) (path.Path, error)
}

type getItemByIDer interface {
	GetItemByID(ctx context.Context, itemID string) (models.Listable, *details.SharePointInfo, error)
}

type getItemser interface {
	GetItems(ctx context.Context, cc api.CallConfig) ([]models.Listable, error)
}

type restoreHandler interface {
	PostLister
	PatchLister
	DeleteLister
	GetLister
	GetListsByCollisionKeyser
}

type PostLister interface {
	PostList(
		ctx context.Context,
		listName string,
		storedList models.Listable,
		errs *fault.Bus,
	) (models.Listable, error)
}

type PatchLister interface {
	PatchList(
		ctx context.Context,
		listID string,
		list models.Listable,
	) (models.Listable, error)
}

type DeleteLister interface {
	DeleteList(
		ctx context.Context,
		listID string,
	) error
}

type GetLister interface {
	GetList(
		ctx context.Context,
		listID string,
	) (models.Listable, *details.SharePointInfo, error)
}

type GetListsByCollisionKeyser interface {
	// GetListsByCollisionKey looks up all lists currently in
	// the site, and returns them in a map[collisionKey]listID.
	// The collision key is displayName of the list
	// which uniquely identifies the list.
	// Collision key checks are used during restore to handle the on-
	// collision restore configurations that cause the list restore to get
	// skipped, replaced, or copied.
	GetListsByCollisionKey(ctx context.Context) (map[string]string, error)
}
