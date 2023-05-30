package exchange

import (
	"context"

	"github.com/alcionai/corso/src/internal/connector/graph"
	"github.com/alcionai/corso/src/pkg/backup/details"
	"github.com/alcionai/corso/src/pkg/fault"
	"github.com/alcionai/corso/src/pkg/path"
	"github.com/alcionai/corso/src/pkg/services/m365/api"
)

// ---------------------------------------------------------------------------
// restore
// ---------------------------------------------------------------------------

type restoreHandler interface {
	itemRestorer
	containerAPI
	newContainerCache(userID string) graph.ContainerResolver
	formatRestoreDestination(
		destinationContainerName string,
		collectionFullPath path.Path,
	) *path.Builder
}

// runs the item restoration (ie: item creation) process
// for a single item, whose summary contents are held in
// the body property.
type itemRestorer interface {
	restore(
		ctx context.Context,
		body []byte,
		userID, destinationID string,
		errs *fault.Bus,
	) (*details.ExchangeInfo, error)
}

// runs the actual graph API post request.
type itemPoster[T any] interface {
	PostItem(
		ctx context.Context,
		userID, dirID string,
		body T,
	) (T, error)
}

// produces structs that interface with the graph/cache_container
// CachedContainer interface.
type containerAPI interface {
	// POSTs the creation of a new container
	CreateContainer(
		ctx context.Context,
		userID, containerName, parentContainerID string,
	) (graph.Container, error)

	// GETs a container by name
	containerSearcher() (containerByNamer, bool)

	// returns either the provided value (assumed to be the root
	// folder for that cache tree), or the default root container
	// (if the category uses a root folder that exists above the
	// restore location path).
	orRootContainer(string) string
}

// searches for a container by name.
// normally, we'd alias the func directly.  The indirection here
// is because not all types comly with GetContainerByName.  We
// identify those that aren't because `containerSearcher()` will
// return (nil, false), in that case.
type containerByNamer interface {
	GetContainerByName(
		ctx context.Context,
		userID, containerName string,
	) (graph.Container, error)
}

// primary interface controller for all per-cateogry restoration behavior.
func restoreHandlers(
	ac api.Client,
) map[path.CategoryType]restoreHandler {
	return map[path.CategoryType]restoreHandler{
		path.ContactsCategory: newContactRestoreHandler(ac),
		path.EmailCategory:    newMailRestoreHandler(ac),
		path.EventsCategory:   newEventRestoreHandler(ac),
	}
}
