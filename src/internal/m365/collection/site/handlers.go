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
	DeleteLister
}

type PostLister interface {
	PostList(
		ctx context.Context,
		listName string,
		storedList models.Listable,
		errs *fault.Bus,
	) (models.Listable, error)
}

type DeleteLister interface {
	DeleteList(
		ctx context.Context,
		listID string,
	) error
}
